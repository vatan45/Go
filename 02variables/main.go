package main

import "fmt"

const LoginToken string = "vatanmalik04" //public

func main() {
	var username string = "vatan"
	fmt.Println(username)

	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)

	fmt.Printf("variable is of type: %T \n", isLoggedin)

	var smallval uint8 = 4
	fmt.Println(smallval)

	fmt.Printf("variable is of type: %T \n", smallval)

	var smallFloat float32 = 255.5645645645646464
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	var largefloat float64 = 255.5645645645646464
	fmt.Println(largefloat)
	fmt.Printf("variable is of type: %T \n", largefloat)

	var whiisin int
	fmt.Println(whiisin)
	fmt.Printf("variable is of type: %T \n", largefloat)

	//implicit type
	var name = "vatan malik"
	fmt.Println(name)

	fmt.Println(LoginToken)

}
