package algos

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	} else if (root.Left == nil && root.Right == nil) && (targetSum == root.Val) {
		return true
	}

	targetSum -= root.Val

	return PathSum(root.Left, targetSum) || PathSum(root.Right, targetSum)
}
