package main

import (
	"fmt"
	"unicode"
)

func PasswordChecker(pw string) bool {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	pw := "11111!ddD"
	if PasswordChecker(pw) {
		fmt.Println("Is Strong password")
	} else {
		fmt.Println("Is Bad Password")
	}
}
