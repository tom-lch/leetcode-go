package PathSumII

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var ans [][]int
func PathSum(root *TreeNode, sum int) [][]int {
	defer func() {
		ans = nil
	}()
	if root == nil {
		return nil
	}
	var arr []int
	addNode(root, sum, arr)
	return ans
}

func addNode (node *TreeNode, sum int, arr []int) {
	if node == nil {
		return
	}
	arr = append(arr, node.Val)
	sum -= node.Val
	if sum == 0 && node.Left==nil && node.Right==nil {
		ans = append(ans, arr)
	}
	a := make([]int, len(arr))
	copy(a, arr)
	addNode(node.Left, sum, a)
	addNode(node.Right, sum, a)
}

