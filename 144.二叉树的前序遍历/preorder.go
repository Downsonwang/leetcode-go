package _44_二叉树的前序遍历

type TreeNode struct {
	   Val int
	    Left *TreeNode
	   Right *TreeNode
	 }
	 // 根左右
func preorderTraversal(root *TreeNode) []int {
	 var res []int
	 var preorder  func(node *TreeNode)

	 preorder = func(node *TreeNode) {
		 if node == nil {
			 return
		 }

		 res = append(res,node.Val)
		 preorder(node.Left)
		 preorder(node.Right)
	 }
	 preorder(root)
	 return res
}