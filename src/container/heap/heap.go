// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heap provides heap operations for any type that implements
// heap.Interface. A heap is a tree with the property that each node is the
// minimum-valued node in its subtree.
//
// The minimum element in the tree is the root, at index 0.
//
// A heap is a common way to implement a priority queue. To build a priority
// queue, implement the Heap interface with the (negative) priority as the
// ordering for the Less method, so Push adds items while Pop removes the
// highest-priority item from the queue. The Examples include such an
// implementation; the file example_pq_test.go has the complete source.
package heap

import "sort"

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// Init has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that Push and Pop in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use heap.Push and heap.Pop.
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init(h Interface) {
	// heapify
	n := h.Len()
	// 从中间开始 直到第0个元素 根节点
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x any) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) any {
	n := h.Len() - 1
	// 将根节点交换到最后
	h.Swap(0, n)
	// 从根节点往下维护堆结构 并排除最后一个节点
	down(h, 0, n)
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h Interface, i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up(h Interface, j int) {
	for {
		// i为j的父节点
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			// i如果越界或者满足堆条件 跳出循环
			break
		}
		// 交换i,j节点
		h.Swap(i, j)
		// 继续对i的父节点检查
		j = i
	}
}

func down(h Interface, i0, n int) bool {
	i := i0
	for {
		// j1是i节点的左子节点
		j1 := 2*i + 1
		// j1是否越界
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		// j2是i节点的右子节点
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			// j2未越界并且j2小(大)于j1
			j = j2 // = 2*i + 2  // right child
		}
		// j为j1,j2之间最小(大)子节点
		if !h.Less(j, i) {
			// j子节点大(小)于等于i父节点
			break
		}
		// 交换父子节点
		h.Swap(i, j)
		// 令i=j,继续循环,保证j节点的子子节点是堆结构
		i = j
	}
	return i > i0
}
