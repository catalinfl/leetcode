package algos

import (
	"strconv"
)

/*
Given two binary strings a and b, return their sum as a binary string.



Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

*/

func AddBinary(a string, b string) string {
	decimalInt, _ := strconv.ParseInt(a, 2, 0)
	decimalInt2, _ := strconv.ParseInt(b, 2, 0)
	numToBin := decimalInt + decimalInt2
	nr := strconv.FormatInt(int64(numToBin), 2)
	return nr
}
