package algos

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	str := ""
	for _, val := range words {
		temp := ""
		for i := len(val) - 1; i >= 0; i-- {
			temp += string(val[i])
		}

		str = strings.Join([]string{str, temp}, " ")
	}

	return strings.Replace(str, " ", "", 1)

}
