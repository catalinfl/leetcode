package algos

func addDigits(num int) int {
	return helper3(num)
}

func helper3(num int) int {
	digits := 0
	sum := 0
	for num != 0 {
		digit := num % 10
		num /= 10
		sum += digit
		digits++
	}
	if digits >= 2 {
		return addDigits(sum)
	} else {
		return sum
	}
}
