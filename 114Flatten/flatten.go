package _14Flatten

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//给你二叉树的根结点 root ，请你将它展开为一个单链表：

//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。
// 输入：root = [1,2,5,3,4,null,6]
//输出：[1,null,2,null,3,null,4,null,5,null,6]


func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	l,r := root.Left,root.Right
	predecessor := l
	if l != nil {
		for predecessor.Right != nil {
			predecessor = predecessor.Right
		}
		predecessor.Right = r
		root.Left = nil
		root.Right = l
	}
	flatten(root.Right)
}
