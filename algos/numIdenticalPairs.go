package algos

func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

func numIdenticalPairs2(nums []int) int {
	var z = make(map[int]int)
	var res int = 0
	for i := 0; i < len(nums); i++ {
		res += z[nums[i]]
		z[nums[i]]++
	}
	return res
}
