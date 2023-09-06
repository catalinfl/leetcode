package algos

import (
	"strings"
)

/*
Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".



Example 1:

Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
Example 2:

Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"

*/

func DefangIPaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}

func DefangIPaddr2(address string) string {
	a := "."
	b := "[.]"

	rep := ""

	for _, s := range address {
		if string(s) == a {
			rep += b
			continue
		}
		rep += string(s)
	}
	return rep
}
