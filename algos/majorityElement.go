package algos

import "math"

func MajorityElement(nums []int) int {
	mapcon := make(map[int]int, len(nums))
	max := math.MinInt
	var maxCount int = 0
	for i := 0; i < len(nums); i++ {
		mapcon[nums[i]]++
		if mapcon[nums[i]] > maxCount {
			max = nums[i]
			maxCount = mapcon[nums[i]]
		}
	}

	if maxCount >= int(len(nums))/2 {
		return max
	}

	return 0
}
