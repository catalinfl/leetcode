package algos

func MinDepth(root *TreeNode) int {
	count := 1
	mindepthhelper(root, count)
	return count
}

func mindepthhelper(root *TreeNode, count int) int {
	left, right := 0, 0

	if root.Left == nil && root.Right == nil {
		return count
	}

	if root.Left != nil {
		left = mindepthhelper(root.Left, count+1)
	}

	if root.Right != nil {
		right = mindepthhelper(root.Right, count+1)
	}

	if left > right {
		return right
	} else {
		return left
	}
}
