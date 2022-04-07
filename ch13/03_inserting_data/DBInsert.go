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
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load environment variables")
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	defer db.Close()

	// NOTE: to prevent sql injection, we use the Prepare method
	insert, err := db.Prepare("INSERT INTO test (id, name) VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(1, "Mustafa") // NOTE: you can insert duplicate data
	if err != nil {
		panic(err)
	} else {
		fmt.Println("the value was successfully inserted")
	}
}
