package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
     treeRoot := root
	 stack := []*TreeNode{}
	 inorder := math.MinInt64
	 for treeRoot != nil || len(stack) > 0 {
		 for root != nil {
			 stack = append(stack,root)
			 root = root.Left
		 }
		 root = stack[len(stack) - 1]
		 stack = stack[:len(stack) - 1]
         if root.Val <= inorder {
			 return false
		 }
		 inorder = root.Val
		 root = root.Right
	 }
	 fmt.Println(stack)
	 return true
}
func main() {
	root := &TreeNode{Val: 2,Left: &TreeNode{Val: 1},Right: &TreeNode{Val: 3}}
	d := IsValidBST(root)
	fmt.Println(d)
}
