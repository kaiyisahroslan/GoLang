package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Basic function call with no return values
	welcome()

	// Functions can return values
	// name := "Amir"
	// check := welcomeByName(name)
	// if check {
	// 	fmt.Printf("Welcome, %v!\n", name)
	// } else {
	// 	fmt.Println(name, ", your name is not in the guest list")
	// }

	// CLI
	name := os.Args[1]
	check := welcomeByName(name)
	if check {
		fmt.Printf("Welcome, %v!\n", name)
	} else {
		fmt.Println(name, ", your name is not in the guest list\n")
	}

	// ... means that give this array the size as per count of value
	scores := []int{10, 20, 30, 40}
	length, total := getSizeAndCount(scores)
	fmt.Println("Length: ", length)
	fmt.Println("Total: ", total)
}

// func getSizeAndCount(s []int) (int, int) {
// 	size := len(s)
// 	sum := 0

// 	for _, v := range s {
// 		sum += v
// 	}

// 	return size, sum // can return more than 1 value
// }

func getSizeAndCount(s []int) (size int, sum int) {
	size = len(s)
	sum = 0

	for _, v := range s {
		sum += v
	}

	return
}

func welcomeByName(nm string) bool {

	result := strings.ContainsAny(nm, "a || e || i || o || u")
	if result {
		return true
	} else {
		return false
	}

	// result1 := strings.Contains(nm, "a")
	// result2 := strings.Contains(nm, "e")
	// result3 := strings.Contains(nm, "i")
	// result4 := strings.Contains(nm, "o")
	// result5 := strings.Contains(nm, "u")

	// if result1 || result2 || result3 || result4 || result5 {
	// 	return true
	// } else {
	// 	return false
	// }
}

func welcome() {
	fmt.Println("Welcome to Uranus!")
}
