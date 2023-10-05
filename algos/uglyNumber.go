package algos

func isUgly(n int) bool {
	if n == 1 {
		return true
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		}
		if n%3 == 0 {
			n /= 3
		}
		if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}

	return true

}
