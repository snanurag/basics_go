package main

import (
	"fmt"
	"github.com/snanurag/basics_go/collections"
)
import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type PQ struct {
	collections.PriorityQueue
}

func (p *PQ) Less(i, j int) bool {
	return p.PriorityQueue[i].Value.(int) < p.PriorityQueue[j].Value.(int)
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := PQ{}
	final := &ListNode{
		Val:  0,
		Next: nil,
	}
	for _, l := range lists {
		for l != nil {
			item := collections.Item{
				Value: l.Val,
			}
			heap.Push(&pq, &item)
			l = l.Next
		}
	}

	head := final
	for pq.Len() > 0 {
		final.Next = &ListNode{
			Val: heap.Pop(&pq).(*collections.Item).Value.(int),
		}
		final = final.Next
	}
	return head.Next
}

func main() {
	var list []*ListNode

	list = append(list, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	})
	list = append(list, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	})

	list = append(list, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	})
	f := mergeKLists(list)
	for f != nil {
		fmt.Println(f.Val)
		f = f.Next

	}
}

/**
class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        head = point = ListNode(0)
        q = PriorityQueue()
        for l in lists:
            if l:
                q.put((l.val, l))
        while not q.empty():
            val, node = q.get()
            point.next = ListNode(val)
            point = point.next
            node = node.next
            if node:
                q.put((node.val, node))
        return head.next
*/
