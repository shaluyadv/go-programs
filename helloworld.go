package main

import (
	"fmt"
)

func main() {

	var myNum1 int = 90

	var myWord1 string = "Hello"

	var myBool bool = true

	fmt.Println(myNum1)
	fmt.Println(myWord1)
	fmt.Println(myBool)

	var a = 10
	var b = "Hello"
	var c = 20
	var d = "World"

	fmt.Println(a, b, c, d)

	var cars = [5]string{"Bolero", "Ferrari", "BMW", "Mercedes", "Scorpio"}
	fmt.Println(cars[4])

	var myNum2 = 60
	var myNum3 = 70
	var z = myNum2 + myNum3

	if myNum2 == myNum3 {

		fmt.Println(100)

		/* if myNum2 will be equal to myNum3 Output would be 100. */

	} else if myNum2 > myNum3 {
		fmt.Println(60)

		/* if myNum2 will be greater than myNum3 Output would be 60. */

	} else {
		fmt.Println(z)
	}

	/* if those above conditions do not match then the output would be myNum2+myNum3=130. */
}
