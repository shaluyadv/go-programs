package main

import (
	"fmt"
)

func main() {

	//Multiply 10 with 5, and print the result.

	var a, b = 10, 5
	var c = a * b
	fmt.Println(c)

	// 0r fmt.Println(10*5)

	//Use the correct arithmetic operator to increase the value of the variable x by 1.

	x := 1
	x++
	fmt.Println(x)

	//Use assignment operators to add 5 to x.

	x += 5
	fmt.Println(x)

}
