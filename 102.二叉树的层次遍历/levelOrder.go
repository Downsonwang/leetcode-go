package _02_二叉树的层次遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
   二叉树的先前遍历(根左右) 记录每一次的二叉树深度 层次递归遍历 每次进入下一层 深度+1 若当前深度大于等于当前res的长度
   给 res 扩容 因为是先序根左右 符合层序遍历的层级次序.

*/
func levelOrder(root *TreeNode) (reply [][]int){
	var leverlOrder func(node *TreeNode,depath int)
	leverlOrder = func(node *TreeNode, depth int) {
		  if node == nil{
			  return
		  }
		if depth >= len(reply) {
			reply = append(reply,[]int{})
		}
		reply[depth] = append(reply[depth],node.Val)

		leverlOrder(node.Left,depth+1)
		leverlOrder(node.Right, depth+1)
		return
	}
	leverlOrder(root,0)
	return

}
