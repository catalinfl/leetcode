package algos

import "strconv"

func hammingDistance(x int, y int) int {
	xConv := strconv.FormatInt(int64(x), 2)
	yConv := strconv.FormatInt(int64(y), 2)

	i := len(xConv) - 1
	j := len(yConv) - 1

	count := 0

	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if xConv[i] != yConv[j] {
			count++
		}
	}

	for ; i >= 0; i-- {
		if xConv[i] == '1' {
			count++
		}
	}

	for ; j >= 0; j-- {
		if yConv[j] == '1' {
			count++
		}
	}

	return count
}
