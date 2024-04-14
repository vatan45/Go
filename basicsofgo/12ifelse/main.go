package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	LoginCount := 23
	var result string
	if LoginCount < 10 {
		result = "regular user"

	} else {
		result = "something else"
	}
	fmt.Println((result))
}
