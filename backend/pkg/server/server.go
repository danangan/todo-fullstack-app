package server

import (
	"app/graphql/generated"
	"app/pkg/db"
	"app/pkg/graphql/directives"
	"app/pkg/graphql/resolvers"
	"app/pkg/middleware"
	_tokenManager "app/pkg/tokenManager"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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

	tokenManager := _tokenManager.NewTokenManager(redisClient)

	gqlConfig := createGqlConfig(db, redisClient, tokenManager)

	authMiddleware := middleware.CreateAuthMiddleware(db, redisClient, tokenManager)

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", authMiddleware(createGqlHandler(gqlConfig)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func createGqlConfig(db *gorm.DB, redisClient *redis.Client, tokenManager *_tokenManager.TokenManager) generated.Config {
	directives := directives.NewDirectiveRoot()
	resolvers := resolvers.NewResolver(db, redisClient, tokenManager)

	return generated.Config{Resolvers: resolvers, Directives: *directives}
}

func createGqlHandler(config generated.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(config))

		gqlServer.ServeHTTP(w, r)
	})
}
