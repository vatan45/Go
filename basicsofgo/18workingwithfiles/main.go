package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "This needs to go in a file - vatanmalik"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filenme string) {
	databytes, err := ioutil.ReadFile(filenme)
	if err != nil {
		panic(err)

	}

	fmt.Println("text data insisde the file is \n", databytes)

}
