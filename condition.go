package main

import (
	"fmt"
)

func main() {
	//Print "Hello World" if x is greater than y.
	var x = 50
	var y = 10

	if x > y {
		fmt.Println("Hello World")

	}

	var a = 50
	var b = 500
	if a == b {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}

	//Print "1" if x is equal to y, print "2" if x is greater than y, otherwise print "3".

	if a == b {
		fmt.Println("True")
	} else if a > b {
		fmt.Println("False")
	} else {
		fmt.Println(3)
	}

}
