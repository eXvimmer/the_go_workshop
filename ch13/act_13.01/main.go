package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	id    int
	name  string
	email string
}

func main() {
	users := []User{
		{id: 1, name: "Mustafa", email: "mustafa@gmail.com"},
		{id: 2, name: "Malena", email: "malena@gmail.com"},
		{id: 3, name: "Maya", email: "maya@gmail.com"},
	}
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

	tableCreate := `
  CREATE TABLE users
  (
      id integer NOT NULL,
      name text COLLATE pg_catalog."default" NOT NULL,
      email text COLLATE pg_catalog."default" NOT NULL,
      CONSTRAINT "Users_pkey" PRIMARY KEY (id)
  )
  WITH (
      OIDS = FALSE
  )
  TABLESPACE pg_default;
  ALTER TABLE users
      OWNER to postgres;
`
	_, err = db.Exec(tableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table created successfully")
	}

	insert, err := db.Prepare("INSERT INTO users (id, name, email) VALUES($1,$2,$3)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	for _, u := range users {
		_, err = insert.Exec(u.id, u.name, u.email)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf(
				"The user with name: %s and email %s was successfully added\n",
				u.name, u.email,
			)
		}
	}

	update, err := db.Prepare("UPDATE users SET email=$1 WHERE id=$2")
	if err != nil {
		panic(err)
	}
	defer update.Close()
	_, err = update.Exec("user@gmail.com", 1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The user's email address is successfully updated")
	}

	delete, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		panic(err)
	}
	defer delete.Close()

	_, err = delete.Exec(2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The user with id of 2 is successfully deleted")
	}
}
