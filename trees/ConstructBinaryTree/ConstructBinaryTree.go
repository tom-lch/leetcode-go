package ConstructBinaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val:postorder[len(postorder)-1]}
	var index int
	for i := 0; i < len(inorder); i++{
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
