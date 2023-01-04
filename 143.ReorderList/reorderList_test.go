package _43_ReorderList

import "testing"

func TestReorderList(t *testing.T) {
	root := &ListNode{Val: 1,Next: &ListNode{Val: 2 ,Next:&ListNode{Val: 3,Next: &ListNode{Val: 4}} },}
	 reorderList(root)

}
