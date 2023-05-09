package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	fmt.Println(h)
	heap.Push(h, 9)
	heap.Push(h, 8)
	heap.Push(h, 7)
	heap.Push(h, 6)
	heap.Push(h, 5)
	heap.Push(h, 4)
	heap.Push(h, 3)
	fmt.Println(h)
	heap.Remove(h, 1)
	fmt.Println(h)
	(*h)[2] = 10
	fmt.Println(h)
	heap.Fix(h, 2)
	fmt.Println(h)

	h = &IntHeap{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	heap.Init(h)
	// 			   |------<->-----|
	// 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
}
