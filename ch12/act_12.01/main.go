package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type budgetCategory string

const (
	autoFuel   budgetCategory = "fuel"
	food       budgetCategory = "food"
	mortgage   budgetCategory = "mortgage"
	repairs    budgetCategory = "repairs"
	insurance  budgetCategory = "insurance"
	utilities  budgetCategory = "utilities"
	retirement budgetCategory = "retirement"
)

var (
	ErrInvalidBudgetCategory = errors.New("budget category not found")
)

type transaction struct {
	id       int
	payee    string
	spent    float64
	category budgetCategory
}

func main() {
	bankFile := flag.String("c", "", "the location of bank transaction csv file")
	logFile := flag.String("l", "", "logging of error")
	flag.Parse()

	if *bankFile == "" {
		fmt.Println("csv file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *logFile == "" {
		fmt.Println("log file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := os.Stat(*bankFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("bank file doesn't exist", *bankFile)
			os.Exit(1)
		}
	}

	_, err = os.Stat(*logFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("log file doesn't exist", *logFile)
			os.Exit(1)
		}
	}

	csvFile, err := os.Open(*bankFile)
	if err != nil {
		fmt.Println("error opening file:", *bankFile)
		os.Exit(1)
	}
	defer csvFile.Close()
	trxs := parseBankFile(csvFile, *logFile)
	for _, trx := range trxs {
		fmt.Printf("%v\n", trx)
	}
}

func parseBankFile(bankTransactions io.Reader, logFile string) []transaction {
	r := csv.NewReader(bankTransactions)
	trxs := []transaction{}
	header := true

	for {
		trx := transaction{}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !header {
			for i, v := range record {
				switch i {
				case 0: // id
					v = strings.TrimSpace(v)
					trx.id, err = strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
				case 1: // payee
					v = strings.TrimSpace(v)
					trx.payee = v
				case 2: // spent
					v = strings.TrimSpace(v)
					trx.spent, err = strconv.ParseFloat(v, 64)
					if err != nil {
						log.Fatal(err)
					}
				case 3: // category
					trx.category, err = convertToBudgetCategory(v)
					if err != nil {
						s := strings.Join(record, ", ")
						writeErrorToLog("error converting csv category column - ", err, s,
							logFile)
					}
				}
			}
			trxs = append(trxs, trx)
		}
		header = false
	}
	return trxs
}

func convertToBudgetCategory(v string) (budgetCategory, error) {
	v = strings.TrimSpace(strings.ToLower(v))
	switch v {
	case "fuel", "gas":
		return autoFuel, nil
	case "food":
		return food, nil
	case "mortgage":
		return mortgage, nil
	case "repairs":
		return repairs, nil
	case "car insurance", "life insurance":
		return insurance, nil
	case "utilities":
		return utilities, nil
	default:
		return "", ErrInvalidBudgetCategory
	}
}

func writeErrorToLog(msg string, err error, data string, logFile string) error {
	msg += "\n" + err.Error() + "\nData: " + data + "\n\n"
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(msg); err != nil {
		return err
	}
	return nil
}
