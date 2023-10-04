package algos

import "fmt"

func Multiply(num1 string, num2 string) string {
	fmt.Println(transformToInt(num1), transformToInt(num2))
	multiply := transformToInt(num1) * transformToInt(num2)
	return transformToString(multiply)
}

func transformToInt(num string) int {
	res := 0
	for i := 0; i < len(num); i++ {
		if string(num[i]) != "0" {
			res = res*10 + convertDigitToInt(num[i])
		} else {
			res = res*10 + 0
		}
	}
	return res
}

func convertDigitToInt(num byte) int {
	switch string(num) {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	}
	return 0
}

func transformToString(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	reversedNum := 0
	for num != 0 {
		digit := num % 10
		reversedNum = reversedNum*10 + digit
		num /= 10
	}

	res := ""
	for reversedNum != 0 {
		digit := reversedNum % 10
		res += convertDigitsToString(digit)
		reversedNum /= 10
	}

	if isNegative {
		res = "-" + res
	}

	if res[len(res)-1] == 52 {
		res += "0"
	}

	return res
}

func convertDigitsToString(num int) string {
	switch num {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	}
	return ""
}
