package _4_inorderTraversal

import "testing"

func TestInorder(t *testing.T) {
	root := &TreeNode{Val: 1,Left: nil,Right: &TreeNode{Val: 2,Left: &TreeNode{Val: 3,Left: nil,Right: nil},Right: nil}}
	reply :=  inorderTraversal(root)
	t.Log("show the results :",reply)

}
