package main

import "fmt"

func main() {

	fmt.Println("welcome to loops in golnag")
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("index is %v and value is %v\n", day)

	}

	roughtvalue := 1
	for roughtvalue < 10 {

		if roughtvalue == 5 {
			roughtvalue++
			continue
		}
		fmt.Println("value is : ", roughtvalue)
		roughtvalue++
	}

}
