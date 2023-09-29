package algos

func IsMonotonic(nums []int) bool {
	var monotonicIncreasing = true
	z := false

	for i := 0; i < len(nums)-1; i++ {

		if !z {
			if nums[i] > nums[i+1] {
				monotonicIncreasing = false
				z = true
			}
			if nums[i] != nums[i+1] {
				z = true
			}
		}

		if nums[i] < nums[i+1] && !monotonicIncreasing {
			return false
		}

		if nums[i] > nums[i+1] && monotonicIncreasing {
			return false
		}
	}

	return true

}
