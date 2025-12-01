package main

import (
	"log"
	"net/http"
	"os"

	"last-go/graphql"
	"last-go/internal/database"
	"last-go/internal/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	// Database configuration
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "biruhken")
	dbPassword := getEnv("DB_PASSWORD", "ayana")
	dbName := getEnv("DB_NAME", "last_go")

	// Connect to database
	db, err := database.NewConnection(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize database
	if err := db.Init(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create GraphQL resolver
	resolver := &graphql.Resolver{
		DB: db,
	}

	// Create GraphQL server
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	// Create router
	router := mux.NewRouter()

	// Add middleware
	router.Use(middleware.AuthMiddleware)

	// GraphQL endpoint
	router.Handle("/query", srv)

	// GraphQL playground
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("GraphQL playground available at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
