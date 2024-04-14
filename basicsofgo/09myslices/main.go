package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to codebase of slices")

	var fruitList = []string{"Apple", "Tomato", "peach"}
	fmt.Printf("types of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "mango", "banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) //3 is inclusive so 1 and 2 only
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 456
	highScore[2] = 356
	highScore[3] = 845
	// highScore[4] = 777 //give an error as size is 4 only
	//to give more elements and incrrease the size of the array we can
	//use the append methof
	//lot of memory is saved as initialy only 4 blocks of memory is allocated adn then we have added mnore memory

	highScore = append(highScore, 555, 666, 777)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from sswlices based o index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
