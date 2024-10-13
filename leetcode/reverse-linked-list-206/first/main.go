package first

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var prev *ListNode
	for cur != nil {
		third := cur
		cur = cur.Next
		third.Next = prev
		prev = third
	}
	return prev
}
