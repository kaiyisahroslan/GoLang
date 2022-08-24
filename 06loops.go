package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	number := 20
	for j := 1; j <= 12; j++ {
		fmt.Printf("%v * %v = %v\n", number, j, number*j)
	}

	// Another way
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("20 x %v = %v", i, i*20)
	// }

	fmt.Println("----------------------------")
	var scores = [8]int{4, 6, 8, 2, 0, 3, 5, 8}

	// Iterate through the scores array
	for index, value := range scores {
		fmt.Printf("Index: %v, value: %v\n", index, value)

	}

	fmt.Println("----------------------------")
	// Iterate through the scores array
	for _, value := range scores {
		fmt.Printf("Value: %v\n", value)

	}
}
