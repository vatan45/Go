package main

import "fmt"

func main() {
	fmt.Println("welcome to the codebase of pointers")

	// var ptr *int

	// fmt.Println("value of ponter is ", ptr)

	mynumber := 23

	var ptr = &mynumber

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value is : ", mynumber)

}
