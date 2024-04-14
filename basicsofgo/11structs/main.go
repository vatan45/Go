package main

import "fmt"

func main() {
	fmt.Println("struct in golang")

	//no inheritance in golang, no super or parent

	vatan := User{"vatan", "vatan@go.dev", true, 19}
	fmt.Println((vatan))
	fmt.Printf("vatans details are : %+v\n", vatan)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
