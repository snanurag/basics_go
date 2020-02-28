// This example demonstrates a priority queue built using the heap interface.
package collections

import (
	"container/heap"
)

type Heap struct {
	PQ PriorityQueue
}

func NewHeap() *Heap {
	return &Heap{PQ: PriorityQueue{}}
}

func (h *Heap) Push(i interface{}) {
	heap.Push(&h.PQ, i)
}

func (h *Heap) Pop() interface{} {
	return heap.Pop(&h.PQ)
}

func (h *Heap) Peek() interface{} {
	return h.PQ.arr[0]
}

func (h *Heap) Len() int {
	return h.PQ.Len()
}

// PriorityQueue implements container/heap interface
type PriorityQueue struct {
	arr []interface{}
}

func (pq PriorityQueue) Len() int { return len(pq.arr) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	switch pq.arr[i].(type) {
	case string:
		return pq.arr[i].(string) > pq.arr[j].(string)
	case int:
		return pq.arr[i].(int) > pq.arr[j].(int)
	case float64:
		return pq.arr[i].(float64) > pq.arr[j].(float64)
	case float32:
		return pq.arr[i].(float32) > pq.arr[j].(float32)
	default:
		return false

	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.arr = append(pq.arr, x)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.arr
	n := len(old)
	val := old[n-1]
	old[n-1] = nil // avoid memory leak
	pq.arr = old[0 : n-1]
	return val
}
