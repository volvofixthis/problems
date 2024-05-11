package first

import (
	"math"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listToNumber(l *ListNode) int {
	number := 0
	lc := l
	pow := 0
	for {
		number += lc.Val * int(math.Pow10(pow))
		if lc.Next == nil {
			break
		}
		lc = lc.Next
		pow += 1
	}
	return number
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	l1c := l1
	l2c := l2
	l3c := l3
	for {
		s := 0
		if l1c != nil {
			s += l1c.Val
			l1c = l1c.Next
		}
		if l2c != nil {
			s += l2c.Val
			l2c = l2c.Next
		}
		s += l3c.Val
		l3c.Next = &ListNode{}
		if s >= 10 {
			l3c.Next.Val = 1
			s = s - 10
		}
		l3c.Val = s
		if l1c == nil && l2c == nil {
			if l3c.Next.Val == 0 {
				l3c.Next = nil
			}
			break
		}
		l3c = l3c.Next
	}
	return l3
}

func numberToList(number int) *ListNode {
	if number == 0 {
		return &ListNode{
			Val: number,
		}
	}
	pow := int(math.Log10(float64(number)))
	if pow == 0 {
		return &ListNode{
			Val: number,
		}
	}
	var l *ListNode
	m := number
	for {
		ml := int(math.Pow10(pow))
		digit := m / ml
		m = m - digit*ml
		lc := &ListNode{
			Val:  digit,
			Next: l,
		}
		l = lc
		if pow == 0 {
			break
		}
		pow -= 1
	}
	return l
}

func strToList(s string) *ListNode {
	var l *ListNode
	parts := strings.Split(s, ",")
	for i := len(parts) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(parts[i])
		lc := &ListNode{
			Val:  v,
			Next: l,
		}
		l = lc
	}
	return l
}
