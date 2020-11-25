package main

import (
	"fmt"
	"sync"
)

//func largestRectangleArea(A []int) int {
//	var max float64
//	st := NewStack()
//	for i, v := range A {
//		for st.Peek() != nil && v < A[st.Peek().(int)] {
//			h := v
//			L := i - st.Peek().(int)
//			max = math.Max(max, float64(h*L))
//			if st.Size == 1 {
//				break
//			}
//			st.Pop()
//		}
//		st.Push(i)
//	}
//
//	if st.Size != 0 {
//		if st.Size == 1 {
//			max = math.Max(float64(A[st.Pop().(int)]), max)
//		} else {
//			max = math.Max(float64(largestRectangleArea(getReverseArray(st))), max)
//		}
//	}
//	return int(max)
//}
//
//func getReverseArray(st *stack) []int {
//	var a []int
//	for st.Size > 0 {
//		a = append(a, st.Pop().(int))
//	}
//	return a
//}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * @input A : Integer array
 *
 * @Output Integer
 */

type pair struct {
	height int
	width  int
}

func largestRectangleArea(A []int) int {
	var stack []pair

	max := 0
	for i := len(A) - 1; i >= 0; i-- {
		length := 0
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if top.height >= A[i] {
				length = top.width + length
				if max < top.height*length {
					max = top.height * length
				}
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, pair{height: A[i], width: length + 1})
	}

	length := 0
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		length += top.width
		if max < top.height*length {
			max = top.height * length
		}

	}
	return max
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type element struct {
	data interface{}
	next *element
}

type stack struct {
	lock *sync.Mutex
	head *element
	Size int
}

func (stk *stack) Push(data interface{}) {
	stk.lock.Lock()

	element := new(element)
	element.data = data
	temp := stk.head
	element.next = temp
	stk.head = element
	stk.Size++

	stk.lock.Unlock()
}

func (stk *stack) Pop() interface{} {
	if stk.head == nil {
		return nil
	}
	stk.lock.Lock()
	r := stk.head.data
	stk.head = stk.head.next
	stk.Size--

	stk.lock.Unlock()

	return r
}

func (stk *stack) Peek() interface{} {
	if stk.head == nil {
		return nil
	}

	stk.lock.Lock()
	r := stk.head.data

	stk.lock.Unlock()

	return r
}

func NewStack() *stack {
	stk := new(stack)
	stk.lock = &sync.Mutex{}

	return stk
}

func main() {

	fmt.Println(largestRectangleArea([]int{1, 3, 2, 3}))

	//fmt.Println(largestRectangleArea([]int{90, 58, 69, 70, 82, 100, 13, 57, 47, 18}))
	//fmt.Println(largestRectangleArea([]int{18, 47, 57, 13, 100, 82, 70, 69, 58, 90}))

}
