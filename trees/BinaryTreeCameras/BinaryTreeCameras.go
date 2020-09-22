package BinaryTreeCameras


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

var ans int
func minCameraCover(root *TreeNode) int {
	defer func(){
		ans = 0
	}()
	if root == nil {
		return 0
	}
	if dfs(root) == 2 {
		return ans+1
	}
	return ans
}


func dfs(node *TreeNode) int {
	if node == nil {
		return 1
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	if left == 2 || right == 2 {
		ans ++
		return 0
	} else if left == 0 || right == 0 {
		return 1
	} else {
		return 2
	}
}

// 使用广度优先遍历出了点问题

func BFSminCameraCover(root *TreeNode) int {
	if root !=nil && root.Left == nil && root.Right == nil {
		return 1
	}
	// 思路一 使用广度优先 层理遍历
	var arr []*TreeNode
	var high int
	var num int
	arr = append(arr, root)
	for len(arr) > 0 {
		for _, node := range arr {
			arr = nil
			if node.Left != nil {
				arr = append(arr, node.Left)
				if high > 1 &&  high % 2 == 0 {
					num++
				}
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
				if high > 1 && high % 2 == 0 {
					num++
				}
			}
		}
		high ++
	}
	if high < 3 {
		return num+1
	}
	if root.Left != nil && root.Right != nil {
		return num+2
	} else if root.Left != nil || root.Right != nil {
		return num+1
	}
	return num
}