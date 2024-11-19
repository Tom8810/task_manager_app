package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"task_manager_app/graph"
	"task_manager_app/internal/db"
	"task_manager_app/internal/env"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	if err := env.Getenv("./internal/env/.env"); err != nil {
		panic(fmt.Errorf("cannnot read environmental variables: %v", err))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := db.ConnectDB()
	if err != nil {
		panic(fmt.Errorf("cannot connect database: %v", err))
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
