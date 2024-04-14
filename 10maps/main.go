package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("Js shorts for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages: ", languages)

	for _, value := range languages {
		fmt.Printf("for key v , value is %v\n", value)
	}
}
