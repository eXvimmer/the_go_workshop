package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	id    int
	name  string
	email string
}

type Messages struct {
	userId  int
	message string
}

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

	connectivity := db.Ping()
	if connectivity != nil {
		panic(connectivity)
	} else {
		fmt.Println("Good to go!")
	}

	tableCreate := `
  CREATE TABLE public.messages
  (
      userid integer NOT NULL,
      message character varying(280) COLLATE pg_catalog."default" NOT NULL
  )
  WITH (
      OIDS = FALSE
  )
  TABLESPACE pg_default;

  ALTER TABLE public.messages
      OWNER to postgres;
  `

	_, err = db.Exec(tableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The messages table created successfully")
	}

	messages := []Messages{
		{userId: 2, message: "Hey, where are you?"},
		{userId: 2, message: "Wanna come over?"},
		{userId: 2, message: "I miss you..."},
		{userId: 3, message: "Did you do it?"},
		{userId: 3, message: "Bro, that's awesome"},
	}

	insertionText := "INSERT INTO messages (userid, message) VALUES ($1,$2)"
	insert, err := db.Prepare(insertionText)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	for _, m := range messages {
		_, err = insert.Exec(m.userId, m.message)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("A new message added for user:", m.userId)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the username to see messages: ")
	username, err := reader.ReadString('\n')
	username = strings.TrimRight(username, "\r\n")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Looking for messages for user:", username)
	}

	queryText := `SELECT users.name, users.email, messages.message
    FROM messages
    INNER JOIN users
    ON users.id = messages.userid
    WHERE users.name LIKE $1
  `
	query, err := db.Prepare(queryText)
	if err != nil {
		panic(err)
	}

	messageRows, err := query.Query(username)
	if err != nil {
		panic(err)
	}
	defer messageRows.Close()

	numberof := 0
	for messageRows.Next() {
		numberof++
	}

	var (
		message string
		email   string
		name    string
	)

	if numberof == 0 {
		fmt.Println("The query returned nothing, no such user:", username)
	} else {
		fmt.Println("There are a total of", numberof, "messages from the user:", username)

		result, err := query.Query(username)
		for result.Next() {
			err = result.Scan(&name, &email, &message)
			if err != nil {
				panic(err)
			}
			fmt.Println(
				"The user:", name, "with email:", email,
				"has sent the following message:", message,
			)
		}
	}
}
