package _4swapnodesinpairs

import "testing"

func TestSwapPairs(t *testing.T) {
	 d := ListNode{Val: 1,Next: &ListNode{Val: 2,Next: &ListNode{Val: 3,Next:&ListNode{Val: 4,Next: nil}}}}
     res := swapPairs(&d)
	 t.Log(res)
}
