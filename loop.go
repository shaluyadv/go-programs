package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 50; i += 5 { //for forwards
		fmt.Println(i)
	}

	for a := 100; a > 0; a -= 10 { //for backwards
		fmt.Println(a)
	}

}

/*i:=0; - Initialize the loop counter (i), and set the start value to 0
  i <= 100; - Continue the loop as long as i is less than or equal to 100
  i+=10 - Increase the loop counter value by 10 for each iteration*/
