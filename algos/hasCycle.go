package algos

import "fmt"

/*

Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.


Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

var head *ListNode = nil

func AddNode(val int) {
	node := &ListNode{val, nil}
	if head == nil {
		head = node
	} else {
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func PrintList() {
	if head == nil {
		fmt.Println("No items in list")
	}
	for n := head; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}

func HasCycle() bool {
	if head == nil || head.Next == nil {
		return false
	}
	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 == nil || p2.Next == nil {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return true
}

func HasCycProb() {

	AddNode(4)
	AddNode(5)
	AddNode(4)
	fmt.Println(HasCycle())

}
