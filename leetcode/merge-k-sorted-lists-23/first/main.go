package first

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Value struct {
	A    int
	Node *ListNode
}

type MergeHeap []Value

func (h MergeHeap) Len() int { return len(h) }
func (h MergeHeap) Less(i, j int) bool {
	return h[i].Node.Val < h[j].Node.Val
}
func (h MergeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MergeHeap) Push(x any) {
	*h = append(*h, x.(Value))
}

func (h *MergeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var h = MergeHeap{}

func mergeKLists(lists []*ListNode) *ListNode {
	h = h[:0]
	heap.Init(&h)
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(&h, Value{Node: lists[i], A: i})
		lists[i] = lists[i].Next
	}
	result := &ListNode{}
	current := result
	for h.Len() != 0 {
		f := heap.Pop(&h).(Value)
		current.Next = f.Node
		current = current.Next
		if lists[f.A] != nil {
			heap.Push(&h, Value{Node: lists[f.A], A: f.A})
			lists[f.A] = lists[f.A].Next
		}
	}
	return result.Next
}
