package _5reversenodesinkgroup

import "testing"

func TestReverseGroup(t *testing.T) {
	heads := &ListNode{Val: 1,Next: &ListNode{Val: 2,Next: &ListNode{Val: 3,Next: nil}}}
	res := reverseGroup(heads,2)
	t.Log(res)
}
