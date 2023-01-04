/*
 * @Descripttion: 876.链表的中间结点  - 快慢指针
 * @Author: downson
 * @Date: 2022-09-22 10:43:37
 * @LastEditTime: 2022-09-22 15:28:37
 */
package middlenode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
