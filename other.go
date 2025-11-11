package main

import "fmt"

func main() {
	for i := 1; i < 50; i++ {
		if i%2 == 0 {
			fmt.Printf("num: %d is even", i)
		} else {
			fmt.Printf("num: %d isn`t even", i)
		}

	}
}
