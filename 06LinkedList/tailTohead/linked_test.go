package tailtohead

import (
	"testing"
)

func TestReversePrint(t *testing.T) {

	i := ReversePrint(&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}})
	t.Log(i)
}
