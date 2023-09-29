package algos

import "math"

func IsBalanced(root *TreeNode) bool {

	countLeft := 0
	countRight := 0

	helper(root, &countLeft, &countRight)

	i := math.Abs(float64(countLeft - countRight))
	if i > 1 {
		return false
	} else {
		return true
	}
}

func helper(root *TreeNode, countLeft *int, countRight *int) {
	if root.Left != nil {
		*countLeft++
		helper(root.Left, countLeft, countRight)
	}
	if root.Right != nil {
		*countRight++
		helper(root.Right, countLeft, countRight)
	}
}
