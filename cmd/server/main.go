package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	projectURL := os.Getenv("SUPABASE_URL")
	db, err := sql.Open("postgres", projectURL)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()
	// fmt.Print(projectURL)
	// rows, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	fmt.Print(err)
	// 	log.Fatal("Error loading qeurerying")
	// }
	// fmt.Print(rows)
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	fmt.Print("passed!")

}
