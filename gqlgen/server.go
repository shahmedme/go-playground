package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/shakilahmmeed/gqlgen-todos/graph"
	"github.com/shakilahmmeed/gqlgen-todos/graph/generated"
	"github.com/shakilahmmeed/gqlgen-todos/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.sqlite3"), &gorm.Config{})

	if err != nil {
		panic("Error while connecting database")
	}

	db.AutoMigrate(&models.Link{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please contact with %s", "Shakil Ahmed")
}
