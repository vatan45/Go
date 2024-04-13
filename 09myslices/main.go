package main

import "fmt"

func main() {
	fmt.Println("welcome to codebase of slices")

	var fruitList = []string{"Apple", "Tomato", "peach"}
	fmt.Printf("types of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "mango", "banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

}
