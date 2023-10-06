package algos

import "strconv"

func hammingWeight(num int32) int {
	numBits := strconv.FormatInt(int64(num), 2)
	count := 0

	for _, bit := range numBits {
		if bit == '1' {
			count++
		}
	}

	return count
}
