package _4_inorderTraversal

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}
func inorderTraversal(root *TreeNode) []int{
	var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res


}
