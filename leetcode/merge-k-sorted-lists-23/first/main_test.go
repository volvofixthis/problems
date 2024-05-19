package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func arrToListNode(a []int) *ListNode {
	l := &ListNode{}
	c := l
	for i := 0; i < len(a); i++ {
		c.Next = &ListNode{Val: a[i]}
		c = c.Next
	}
	return l.Next
}

func arrsToListNodes(a [][]int) []*ListNode {
	r := []*ListNode{}
	for i := 0; i < len(a); i++ {
		r = append(r, arrToListNode(a[i]))
	}
	return r
}

func listNodeToArr(l *ListNode) []int {
	r := []int{}
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r
}

func TestPointer(t *testing.T) {
	l := &ListNode{}
	c := l
	*l = ListNode{Val: 2}
	assert.Equal(t, c.Val, l.Val, "should be equal")
}

func TestConvertors(t *testing.T) {
	tests := []struct {
		name string
		i1   []int
		o2   []int
	}{
		{"first", []int{1, 4, 5}, []int{1, 4, 5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := arrToListNode(test.i1)
			assert.Equal(t, test.o2, listNodeToArr(l), "should be equal")
		})
	}
}

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   [][]int
		o2   []int
	}{
		{"first", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"second", [][]int{}, []int{}},
		{"third", [][]int{{}}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o2, listNodeToArr(mergeKLists(arrsToListNodes(test.i1))), "should be equal")
		})
	}
}
