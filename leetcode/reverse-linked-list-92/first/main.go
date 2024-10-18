package first

import (
	"github.com/volvofixthis/problems/utils/linkedlist"
)

func reverseBetween(head *linkedlist.ListNode, left int, right int) *linkedlist.ListNode {
	cur := head
	var prev *linkedlist.ListNode
	for left > 1 {
		prev = cur
		cur = cur.Next
		left -= 1
		right -= 1
	}
	con := prev
	tail := cur
	for right > 0 {
		third := cur
		cur = cur.Next
		third.Next = prev
		prev = third
		right -= 1
	}
	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}
	tail.Next = cur
	return head
}
