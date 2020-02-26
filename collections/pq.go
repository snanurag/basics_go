// This example demonstrates a priority queue built using the heap interface.
package collections

// An Item is something we manage in a priority queue.
type Item struct {
	Value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	switch pq[i].Value.(type) {
	case string:
		return pq[i].Value.(string) > pq[j].Value.(string)
	case int:
		return pq[i].Value.(int) > pq[j].Value.(int)
	case float32:
		return pq[i].Value.(float32) > pq[j].Value.(float32)
	case float64:
		return pq[i].Value.(float64) > pq[j].Value.(float64)
	default:
		return false

	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//// This example creates a PriorityQueue with some items, adds and manipulates an item,
//// and then removes the items in priority order.
//func main() {
//	// Create a priority queue, put the items in it, and
//	// establish the priority queue (heap) invariants.
//	pq := PriorityQueue{}
//
//	heap.Push(&pq, &Item{
//		Value: "orange",
//	})
//
//	heap.Push(&pq, &Item{
//		Value: "apple",
//	})
//	heap.Push(&pq, &Item{
//		Value: "pear",
//	})
//	heap.Push(&pq, &Item{
//		Value: "banana",
//	})
//	// Take the items out; they arrive in decreasing priority order.
//	for pq.Len() > 0 {
//		item := heap.Pop(&pq).(*Item)
//		fmt.Printf("%.2d:%s ", item.priority, item.Value)
//	}
//}
