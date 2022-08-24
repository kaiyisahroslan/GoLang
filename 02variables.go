package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Declare the variable
	// var variableName dataType
	// Declare variable score
	// you cannot use a variable until it is declared!

	// integers
	var score int
	// Assign value 8 to variable score
	// Initialize variable score with value 8
	score = 8
	score = score + 1

	// If you declare the variable, need to use
	fmt.Println("Go is strictly typed language!")
	fmt.Println(score)

	// Check the datatype of variable, use reflect package
	fmt.Println(reflect.TypeOf(score).Kind())

	// string variable
	var message string
	message = "Go lang makes beginner developers good developers!"

	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))

	// float values
	var lifeline float32
	lifeline = 44.44

	fmt.Println(lifeline)
	fmt.Println(reflect.TypeOf(lifeline))

	// boolean values
	var check bool
	check = true

	fmt.Println(check)
	fmt.Println(reflect.TypeOf(check))

	// auto type inference
	// use := for auto type inference
	// life is not yet declared
	// then use :=
	fmt.Println("Using auto type inference!")
	life := 90
	fmt.Println(life)
	fmt.Println(reflect.TypeOf(life))

	myMessage := "Hello, how are you?"
	fmt.Println(myMessage)
	fmt.Println(reflect.TypeOf(myMessage))

	age := 71.71
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(age))

}
