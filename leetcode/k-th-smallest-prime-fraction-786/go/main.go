package main

import "container/heap"

type Fraction struct {
	n int
	d int
}

type FractionHeap []Fraction

func (h FractionHeap) Len() int { return len(h) }
func (h FractionHeap) Less(i, j int) bool {
	return h[i].n*h[j].d-h[j].n*h[i].d < 0
}
func (h FractionHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *FractionHeap) Push(x any) {
	*h = append(*h, x.(Fraction))
}

func (h *FractionHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	h := &FractionHeap{}
	heap.Init(h)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			heap.Push(h, Fraction{arr[i], arr[j]})
		}
	}
	var f Fraction
	for i := 0; i < k; i++ {
		f = heap.Pop(h).(Fraction)
	}
	return []int{f.n, f.d}
}
