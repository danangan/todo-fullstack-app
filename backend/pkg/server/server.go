package server

import (
	"app/graphql/generated"
	"app/pkg/db"
	"app/pkg/graphql/directives"
	"app/pkg/graphql/resolvers"
	"app/pkg/middleware"
	_todoService "app/pkg/todoService"
	_tokenService "app/pkg/tokenService"
	_userService "app/pkg/userService"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/redis/go-redis/v9"
)

var defaultPort string = "8080"

func CreateServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := db.CreateDBConnection()

	if err != nil {
		log.Fatalf("Failed to create DB connection, error: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if redisClient == nil {
		log.Panicln("failed to start redis client")
	}

	// Service initialization
	tokenService := _tokenService.New(redisClient)
	userService := _userService.New(db)
	todoService := _todoService.New(db)

	gqlConfig := createGqlConfig(tokenService, userService, todoService)

	authMiddleware := middleware.CreateAuthMiddleware(tokenService, userService)

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", authMiddleware(createGqlHandler(gqlConfig)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func createGqlConfig(tokenManager *_tokenService.TokenManager, userService *_userService.UserService, todoService *_todoService.TodoService) generated.Config {
	directives := directives.NewDirectiveRoot()
	resolvers := resolvers.NewResolver(tokenManager, userService, todoService)

	return generated.Config{Resolvers: resolvers, Directives: *directives}
}

func createGqlHandler(config generated.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(config))

		gqlServer.ServeHTTP(w, r)
	})
}
