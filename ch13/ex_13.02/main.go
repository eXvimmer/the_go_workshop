package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
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
		number    int64
		prop      string
		primeSum  int64
		newNumber int64
	)

	allTheNumbers := "SELECT * FROM number"
	nums, err := db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	defer nums.Close()

	primeSum = 0
	result, err := nums.Query()
	if err != nil {
		panic(err)
	}

	fmt.Println("the list of prime numbers:")
	for result.Next() {
		err := result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		if big.NewInt(number).ProbablyPrime(0) {
			primeSum += number
			fmt.Print(" ", number)
		}
	}
	fmt.Println("\nThe total sum of prime numbers in this range is:", primeSum)

	remove := "DELETE FROM number WHERE property=$1"
	removeResult, err := db.Exec(remove, "even")
	if err != nil {
		panic(err)
	}

	modifiedRecords, err := removeResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("The number of rows removed:", modifiedRecords)
	fmt.Println("Updating numbers...")

	update := "UPDATE Number SET number=$1 WHERE number=$2 AND property=$3"
	result, err = nums.Query()
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		newNumber = number + primeSum
		_, err = db.Exec(update, newNumber, number, prop)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("the execution is now complete...")
}
