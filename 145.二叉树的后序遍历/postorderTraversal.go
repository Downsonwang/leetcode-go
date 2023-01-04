package _45_二叉树的后序遍历

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 后序遍历: 左右根
func postorderTraversal(root *TreeNode) []int {
	var res []int
    var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)

		postorder(node.Right)
		res = append(res,node.Val)
	}
	postorder(root)
	return res
}
