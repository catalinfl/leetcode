package main

import (
	"fmt"
)

func main() {
	var i = []int{1, 2, 3, 4}
	fmt.Println(i[len(i)/2])
}
