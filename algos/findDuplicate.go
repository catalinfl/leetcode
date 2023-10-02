package algos

func findDuplicate(nums []int) int {
	r := make(map[int]uint8, len(nums)/2)
	res := 0

	for _, num := range nums {
		r[num]++
		if r[num] == 2 {
			res = num
		}
	}

	return res
}
