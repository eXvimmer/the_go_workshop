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

	var prop string

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

	tableCreate := `
    CREATE TABLE number (
      number integer NOT NULL,
      property text COLLATE pg_catalog."default" NOT NULL
    )
    WITH (
      OIDS = FALSE
    )
    TABLESPACE pg_default;
    ALTER TABLE number
    OWNER TO postgres;
  `

	_, err = db.Exec(tableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("the table called number was created successfully")
	}

	insert, err := db.Prepare("INSERT INTO number VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "even"
		} else {
			prop = "odd"
		}
		_, err := insert.Exec(i, prop)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("The number", i, "is:", prop)
		}
	}
	insert.Close()
}
