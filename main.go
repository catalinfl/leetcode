package main

import (
	"fmt"
)

func main() {
	var z []int = []int{1, 2, 3, 4, 5}
	fmt.Println(len(z))
	z = z[len(z)-1:]
	fmt.Println(z)
}
