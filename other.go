package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		default:
			hasSpecial = true
		}
	}
	score := 0
	if len(password) >= 8 {
		score++
	}
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasDigit {
		score++
	}
	if hasSpecial {
		score++
	}
	fmt.Printf("Password lvl: %d/5", score)
}
