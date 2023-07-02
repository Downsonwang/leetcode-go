package _5reverseKGroup

import "testing"

func TestRevereKGroup(t *testing.T) {
	arr := ListNode{Val: 1,Next: &ListNode{Val: 2,Next: &ListNode{Val: 3,Next: &ListNode{Val: 4,Next: &ListNode{Val: 5,Next: nil}}}}}
	res := reverseKGroup(&arr,2)
	t.Log(res)
}
