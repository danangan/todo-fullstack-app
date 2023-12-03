package server

import (
	"app/graphql/generated"
	"app/pkg/db"
	"app/pkg/graphql/directives"
	"app/pkg/graphql/resolvers"
	"app/pkg/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	gqlConfig := createGqlConfig(db)

	authMiddleware := middleware.CreateAuthMiddleware(db)

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", authMiddleware(createGqlHandler(gqlConfig)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func createGqlConfig(db *gorm.DB) generated.Config {
	directives := directives.NewDirectiveRoot()
	resolvers := resolvers.NewResolver(db)

	return generated.Config{Resolvers: resolvers, Directives: *directives}
}

func createGqlHandler(config generated.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(config))

		gqlServer.ServeHTTP(w, r)
	})
}
