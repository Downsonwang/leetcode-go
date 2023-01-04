/*
 * @Descripttion: 206.反转链表
 * @Author: downson
 * @Date: 2022-09-16 17:02:45
 * @LastEditTime: 2022-09-19 09:13:33
 */
package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	var curr *ListNode = head
	for head != nil {
		var next *ListNode = curr.Next //
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
