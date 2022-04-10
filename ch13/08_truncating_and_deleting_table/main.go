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

	emptyTable, err := db.Exec("TRUNCATE TABLE test")
	if err != nil {
		panic(err)
	}
	rows, err := emptyTable.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d rows affected\n", rows)

	_, err = db.Exec("DROP TABLE test")
	if err != nil {
		panic(err)
	}
	fmt.Println("dropped the test database")
}
