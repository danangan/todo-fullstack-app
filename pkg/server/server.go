package server

import (
	"app/graph/generated"
	appContext "app/pkg/app-context"
	"app/pkg/db"
	"app/pkg/db/models"
	"app/pkg/resolvers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	d := generated.DirectiveRoot{
		Authenticated: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			val := ctx.Value(appContext.CurrentUserKey)

			currentUser := val.(*models.User)
			fmt.Println(currentUser)
			return nil, nil
		},
	}

	c := generated.Config{Resolvers: &resolvers.Resolver{Db: db}, Directives: d}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	contextProviderHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var currentUser *models.User
		token := r.Header.Get("TOKEN")

		if token != "" {
			currentUser = &models.User{
				FirstName: "Danang",
				LastName:  "Nur",
				Email:     "some-email",
			}
		}

		ctx := context.WithValue(r.Context(), appContext.CurrentUserKey, currentUser)

		test := handler.NewDefaultServer(generated.NewExecutableSchema(c))

		test.ServeHTTP(w, r.WithContext(ctx))
	})

	http.Handle("/query", contextProviderHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
