/*
 * @Descripttion: 19.删除链表的倒数第N个结点
 * @Author: downson
 * @Date: 2022-09-22 08:39:03
 * @LastEditTime: 2022-09-22 09:03:38
 */
package removenthfromend

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy

	for i := 0; i < n; i++ {
		head = head.Next
	}
	for head != nil {
		pre, head = pre.Next, head.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next

}
