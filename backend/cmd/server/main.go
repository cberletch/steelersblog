package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cberletch/steelersagg/backend/pkg/tags"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "steelersagg"
)

func main() {

	dbUser := os.Getenv("mysql_user")
	dbPassword := os.Getenv("mysql_password")

	// Create the database connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")

	tagsRepo := tags.NewRepository(db)
	tagsHandler := tags.NewHandler(tagsRepo)

	// Define routes and handlers
	http.HandleFunc("/api/tags", tagsHandler.CreateTag)

	// Start the web server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
