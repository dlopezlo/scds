package utils

import (
	"fmt"
	"log"
	"regexp"
)

func GetUsername(username string) (string, error) {
	// expressió regular de noms d'usuari
	usernameRegex := "^[a-zA-Z][a-z,A-Z,\\-,_,0-9,.]{3,29}$"
	regex, err := regexp.Compile(usernameRegex)
	if err != nil {
		log.Println("Error compiling regular expression", err)
		return "", err
	}

	if !regex.MatchString(username) {
		log.Println("Invalid username", err)
		return "", fmt.Errorf("Invalid username: %s", username)
	}

	// usuari validat
	return username, nil
}
