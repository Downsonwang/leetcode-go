package _4_二叉树的中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 树的特点 来遍历  中序 就是  左 根 右 . 定义一个 order 来 接纳 左树 / 右树    -> 递归方法
func inorderTraversal(root *TreeNode) (res []int) {
	/*
	var result = make([]int, 0)
	fmt.Println(result)
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		result = append(result,root.Val)
	}
	if root.Left == nil && root.Right != nil {

		result = append(result, root.Val,root.Right.Val)
		if root.Right.Left != nil && root.Right.Right == nil {

			result = append(result, root.Right.Left.Val, root.Right.Val)
		}

	}
	if root.Left != nil && root.Right == nil {

		result = append(result, root.Left.Val, root.Val)
	}
	return result
  */
	// 递归
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil{
			return
		}  //递归到节点为nil ,终止.
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}


// 迭代  模拟 进出stack
func inorderDiedai(root *TreeNode) ( res []int){
    stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res,root.Val)
		root = root.Right
	}
	return
}