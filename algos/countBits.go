package algos

import "strconv"

func CountBits(n int) []int {

	var array []int

	for i := 0; i < n+1; i++ {

		bitConv := strconv.FormatInt(int64(i), 2)
		array = append(array, countBitsHelper(bitConv))
	}
	return array
}

func countBitsHelper(n string) int {
	counter := 0
	for i := 0; i < len(n); i++ {
		if n[i] == '1' {
			counter++
		}
	}
	return counter
}
