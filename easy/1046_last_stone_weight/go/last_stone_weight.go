package leetcode

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	for _, v := range stones {
		heap.Push(h, v)
	}

	len := len(stones)
	for len > 1 {
		s1 := heap.Pop(h)
		s2 := heap.Pop(h)
		if s1 == s2 {
			len = len - 2
		} else {
			s1 = s1.(int) - s2.(int)
			heap.Push(h, s1)
			len--
		}
	}

	if len == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	// return true if i is larger than j so
	// that the heap effectively becomes a
	// max heap (as opposed to the default min heap)
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
