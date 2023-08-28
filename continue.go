package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 15; i += 1 {
		if i == 7 {

			continue
		}

		fmt.Println(i)

	}
}
