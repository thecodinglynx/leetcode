package leetcode

import (
	"container/heap"
	"strings"
)

/*
    1. store count of words in hashmap
	2. iterate over map and populate max heap (custom comparator)
	3. get values from heap which will be in correct alphabetical order
*/
func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	h := &IntHeap{}
	res := make([]string, 0)

	for _, word := range words {
		if v, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word] = v + 1
		}
	}

	for key, val := range m {
		heap.Push(h, Value{name: key, count: val})
	}

	for i := 0; i < k; i++ {
		v := heap.Pop(h).(Value)
		res = append(res, v.name)
	}

	return res
}

type IntHeap []Value

type Value struct {
	name  string
	count int
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return strings.Compare(h[i].name, h[j].name) < 0
	}
	return h[i].count > h[j].count
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Value))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
