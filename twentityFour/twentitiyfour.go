/*
 * @Descripttion: 24 两两交换链表中的节点
 * @Author:
 * @Date: 2022-06-04 17:25:31
 * @LastEditTime: 2022-06-04 17:49:08
 */

package twentityfour

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入: head = [1,2,3,4]  输出:[2,1,4,3]
// head =[], []
// head = [1] [1]

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}

	dummy.Next = head //虚拟头结点
	prev := dummy

	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next

		next.Next = head
		prev.Next = next

		prev = head
		head = head.Next
	}
	return dummy.Next
}
