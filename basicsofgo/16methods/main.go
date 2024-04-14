package main

import "fmt"

func main() {
	fmt.Println("struct in golang")

	//no inheritance in golang, no super or parent

	vatan := User{"vatan", "vatan@go.dev", true, 19}
	fmt.Println((vatan))
	fmt.Printf("vatans details are : %+v\n", vatan)
	vatan.GetStatus()
	vatan.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "vatan45@go.dev"
	fmt.Println("Email of this user is : ", u.Email)

}
