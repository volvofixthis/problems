package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrToListNode(a []int) *ListNode {
	l := &ListNode{}
	c := l
	for i := 0; i < len(a); i++ {
		c.Next = &ListNode{Val: a[i]}
		c = c.Next
	}
	return l.Next
}

func ListNodeToArr(l *ListNode) []int {
	r := []int{}
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r
}
