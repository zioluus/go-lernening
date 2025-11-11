package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Print("Age: ")
	fmt.Scan(&age)

	year := 2025 - age
	fmt.Printf("%s your year of birth is: %d", name, year)
}
