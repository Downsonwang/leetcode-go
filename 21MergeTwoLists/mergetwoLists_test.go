/*
 * @Descripttion: 21合并两个有序链表
 * @Author:
 * @Date: 2022-09-20 08:52:37
 * @LastEditTime: 2022-09-20 09:39:07
 */
package mergetwolists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	final := mergeTwoLists(l1, l2)
	t.Log(final)
}
