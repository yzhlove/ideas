package main

import (
	"container/heap"
	"fmt"
)

func main() {

	hp := &IntHeap{}
	heap.Init(hp)
	heap.Push(hp, 1)
	heap.Push(hp, 4)
	heap.Push(hp, 7)
	heap.Push(hp, 6)
	heap.Push(hp, 9)
	heap.Push(hp, 9)
	heap.Push(hp, 8)
	heap.Push(hp, 2)
	heap.Push(hp, 5)

	for hp.Len() > 0 {
		fmt.Println(heap.Pop(hp))
	}

}

type IntHeap []int

func (heap IntHeap) Len() int {
	return len(heap)
}

func (heap IntHeap) Less(i, j int) bool {
	return heap[i] < heap[j]
}

func (heap IntHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *IntHeap) Push(x any) {
	*heap = append(*heap, x.(int))
}

func (heap *IntHeap) Pop() any {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[0 : n-1]
	return x
}
