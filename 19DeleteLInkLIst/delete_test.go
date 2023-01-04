/*
 * @Descripttion:
 * @Author:
 * @Date: 2022-06-01 10:20:45
 * @LastEditTime: 2022-06-01 10:27:50
 */
package deletelinklist

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	lis := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := RemoveNthFromEnd(lis, 2)
	t.Fatal(result)
}
