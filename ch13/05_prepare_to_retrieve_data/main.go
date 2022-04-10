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
		fmt.Println("connected to database")
	}

	defer db.Close()

	var (
		id   int = 2
		name string
	)

	query := "SELECT name FROM test WHERE id=$1"
	queryRow, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer queryRow.Close()

	err = queryRow.QueryRow(id).Scan(&name)
	if err != nil {
		panic(err)
	}

	fmt.Println("id:", id, "name:", name)
}
