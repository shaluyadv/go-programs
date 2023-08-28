package main

import (
	"fmt"
)

func main() {
	adj := [2]string{"Big", "Tasty"}
	fruits := [3]string{"Orange", "Apple", "Mango"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}
