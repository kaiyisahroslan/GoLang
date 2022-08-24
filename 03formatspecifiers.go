package main

import "fmt"

func main() {

	score := 94
	highScore := 99
	// Println() will always print the newline at the end
	// of the statement automatically
	fmt.Println(score)
	fmt.Println("My score is ", score)

	// Formar specifiers
	// use Printf(). cannot use Println()

	// Printf() will not print the newline at the end
	// of the statement automatically
	// use \n to print the new line

	fmt.Printf("My score is %v \n", score)
	fmt.Printf("My score is %v. \n", score)
	fmt.Printf("My score is '%v' \n", score)
	fmt.Printf("My score is: %v \n", score)

	// use multiple format specifiers in the single statement
	fmt.Printf("My score is %v. My highest score till date is %v. \n", score, score)
	fmt.Printf("My score is %v. Highest score till date is %v.\n", score, highScore)

	// %T -> datatype
	fmt.Printf("Value stored in score variable is %v. Datatype of score variabe is %T\n", score, score)

	// working with space in format specifiers
	fmt.Printf("Score is %10v.\n", score) // right align
	fmt.Printf("Score is %-10v.", score)  // left align

}
