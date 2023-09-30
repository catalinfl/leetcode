package algos

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	return helper2(root, count)
}

func helper2(root *TreeNode, count int) int {
	left, right := 0, 0

	if root.Left == nil && root.Right == nil {
		return count
	}

	if root.Left != nil {
		left = helper2(root.Left, count+1)
	}

	if root.Right != nil {
		right = helper2(root.Right, count+1)
	}

	if left > right {
		return left
	} else {
		return right
	}

}
