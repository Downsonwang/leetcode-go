package _4_二叉树的中序遍历

import (
	"testing"
)

func TestInorder(t *testing.T) {
    root := &TreeNode{Val: 1,Left: &TreeNode{Val: 2,Left: nil,Right: nil},Right: &TreeNode{Val: 3,Left: nil,Right: nil}}
	reply :=  inorderTraversal(root)
	 res :=   inorderDiedai(root)
	t.Log("show the results :",reply)
	 t.Log(res)
}
