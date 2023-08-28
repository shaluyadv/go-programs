package main

import (
	"fmt"
)

func main() {

	//Create an array of type string called cars.

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	//Print the value of the second element in the cars array.
	fmt.Println(cars[1])

	//Change the value from "Volvo" to "Opel", in the cars array.

	cars[0] = "Opel"
	fmt.Println(cars)

}
