package tailtohead

type ListNode struct {
	Val  int
	Next *ListNode
}

// 从尾部到头打印链表 尾部填充
func ReversePrint(head *ListNode) []int {

	ans := make([]int, 10000)
	var c = 9999
	for head != nil {
		ans[c] = head.Val
		c--
		head = head.Next
	}
	return ans[c+1:]
}
