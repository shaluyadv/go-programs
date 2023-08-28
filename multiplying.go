package main

import "fmt"

func main() {

	var num1 int = 0
	var num2 int = 0
	var mult int = 0

	fmt.Println("Enter number1: ")
	fmt.Scanf("%d", &num1)

	fmt.Println("Enter number2: ")
	fmt.Scanf("%d", &num2)

	mult = 0
	for count := 1; count <= num2; count++ {
		mult = mult + num1
	}

	fmt.Println("Multiplication is:%d", mult)
}
