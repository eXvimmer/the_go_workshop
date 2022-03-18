package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language  string
	territory string
}

func getLocales() map[locale]struct{} {
	supportedLanguages := make(map[locale]struct{}, 5)
	supportedLanguages[locale{"en", "US"}] = struct{}{}
	supportedLanguages[locale{"fr", "FR"}] = struct{}{}
	supportedLanguages[locale{"en", "CN"}] = struct{}{}
	supportedLanguages[locale{"fr", "CN"}] = struct{}{}
	supportedLanguages[locale{"ru", "RU"}] = struct{}{}

	return supportedLanguages
}

func localeExists(l locale) bool {
	_, ok := getLocales()[l]
	return ok
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error, no locale passed")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("error, invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := locale{
		language:  localeParts[0],
		territory: localeParts[1],
	}

	if !localeExists(passedLocale) {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Locale passed is supported")
}
