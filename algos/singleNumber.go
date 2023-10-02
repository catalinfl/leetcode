package algos

func singleNumber(nums []int) int {
	m := make(map[int]int8, len(nums)/2)

	res := 0

	for _, n := range nums {
		m[n]++
		if m[n] == 1 {
			res += n
		} else {
			res -= n
		}
	}

	return res

}
