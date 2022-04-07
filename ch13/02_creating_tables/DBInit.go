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
		log.Fatal("cannot connect to database")
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	defer db.Close()

	DBCreate := `
    CREATE TABLE public.test
    (
      id integer,
      name character varying COLLATE pg_catalog."default"
    )
    WITH (
      OIDS = FALSE
    )
    TABLESPACE pg_default;
    ALTER TABLE public.test
      OWNER to postgres;
	`
	_, err = db.Exec(DBCreate)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("the table was created successfully")
	}
}
