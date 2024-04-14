package main

import "fmt"

func main() {
	fmt.Println("welcome to the codebase of array")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "peach"

	fmt.Println("fruit list is : ", fruitlist)
	fmt.Println("fruit list is : ", len(fruitlist))
}
