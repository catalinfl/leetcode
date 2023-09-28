package algos

import (
	"sort"
)

func SortArrayByParity(nums []int) []int {
	var returnVec []int
	var evenVec []int

	sort.Ints(nums)

	for _, val := range nums {
		if val%2 == 0 {
			returnVec = append(returnVec, val)
		} else {
			evenVec = append(evenVec, val)
		}
	}

	returnVec = append(returnVec, evenVec...)

	return returnVec
}
