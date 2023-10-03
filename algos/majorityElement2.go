package algos

func majorityElement(nums []int) []int {
	var mapelement = make(map[int]int)
	var elements []int
	var maxToBeWrite = len(nums) / 3
	for i := 0; i < len(nums); i++ {
		mapelement[nums[i]]++
		if (mapelement[nums[i]] >= maxToBeWrite) && (!containsValue(elements, nums[i])) {
			elements = append(elements, nums[i])
		}
	}

	return elements
}

func containsValue(nums []int, number int) bool {
	for _, num := range nums {
		if number == num {
			return true
		}
	}

	return false
}
