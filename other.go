package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Podak imie: ")
	fmt.Scan(&name)

	fmt.Print("Podaj wiek: ")
	fmt.Scan(&age)

	rok_ur := 2025 - age
	fmt.Printf("%s twój rok urodzenia to: %d", name, rok_ur)
}
