package algos

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if (p.Val != q.Val) || (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}