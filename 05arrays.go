// array -> collection of the values of same datatype

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Declare array
	var scores = [8]int{4, 6, 8, 2, 0, 3, 5, 8}
	fmt.Println(scores)

	var friends = [3]string{"aoob", "klia", "cwg"}
	fmt.Println(friends)

	var specificPosition = [9]int{0: 7, 4: 22, 8: 18} // index 0 is 7, index 4 is 22, index 8 is 18
	fmt.Println(specificPosition)

	fmt.Printf("Length of array specificPoisition is %v\n", len(specificPosition))
	fmt.Printf("Data type of array specificPoisition is %T\n", specificPosition)
	fmt.Println("Data type of array specificPoisition is ", reflect.TypeOf(specificPosition).Kind())
}
