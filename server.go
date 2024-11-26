package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"task_manager_app/graph"
	"task_manager_app/internal/db"
	"task_manager_app/internal/env"
	"time"

	"github.com/rs/cors"

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

	// CORS設定
	corsMiddleware := cors.New(cors.Options{
		// 許可するオリジン
		AllowedOrigins: []string{
			"http://localhost:3000",     
			"http://127.0.0.1:3000",
			"http://localhost:8080",     
			"http://127.0.0.1:8080",
		},
		// 許可するHTTPメソッド
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet, 
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch, 
			http.MethodDelete,
			http.MethodOptions,
		},
		// 許可するヘッダー
		AllowedHeaders: []string{
			"*",             
			"Content-Type", 
			"Authorization",
			"Accept",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge: int(30 * time.Minute),
		Debug: true,
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	handler := corsMiddleware.Handler(mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
