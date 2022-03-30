package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	ErrInvalidSSNNumbers    = errors.New("ssn has non-numeric digits")
	ErrInvalidSSNLength     = errors.New("ssn is not nine characters long")
	ErrInvalidSSNPrefix     = errors.New("ssn has three zeros as a prefix")
	ErrInvalidSSNDigitPlace = errors.New("ssn starts with a 9 requires 7 or 9 in the fourth place")
)

func validLength(ssn string) error {
	ssn = strings.TrimSpace(ssn)
	if len(ssn) != 9 {
		return fmt.Errorf("the value of %s caused an error: %v\n",
			ssn, ErrInvalidSSNLength)
	}
	return nil
}

func isNumber(ssn string) error {
	_, err := strconv.Atoi(ssn)
	if err != nil {
		return fmt.Errorf("the value of %s caused an error: %v\n",
			ssn, ErrInvalidSSNNumbers)
	}
	return nil
}

func isPrefixValid(ssn string) error {
	if strings.HasPrefix(ssn, "000") {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNPrefix)
	}
	return nil
}

func validDigitPlace(ssn string) error {
	if string(ssn[0]) == "9" && (string(ssn[3]) != "9" && string(ssn[3]) != "7") {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNDigitPlace)
	}
	return nil
}

func main() {
	ssn := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333",
		"087-65-4321", "123-45-zzzz"}

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	for _, v := range ssn {
		err := validLength(v)
		if err != nil {
			log.Println(err)
		}
		err = isNumber(v)
		if err != nil {
			log.Println(err)
		}
		err = isPrefixValid(v)
		if err != nil {
			log.Println(err)
		}
		err = validDigitPlace(v)
		if err != nil {
			log.Println(err)
		}
	}
}
