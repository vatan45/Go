package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greeter()
	greeterTwo()
	result := adder(3, 5)
	fmt.Println("result is: ", result)

	proResult := proAdder(2, 5, 7, 8)
	fmt.Println("pro result is: ", proResult)

}

func greeterTwo() {
	fmt.Println("another method")
}

func greeter() {
	fmt.Println("Namstey from golnag")

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val

	}
	return total
}
