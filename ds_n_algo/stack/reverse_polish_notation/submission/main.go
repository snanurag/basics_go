package main

//https://www.interviewbit.com/problems/evaluate-expression/

import (
	"strconv"
	"sync"
)

func evalRPN(A []string) int {
	st := New()

	for i := len(A) - 1; i >= 0; i-- {
		v1 := A[i]
		insertNVal(v1, st)
	}
	val, _ := strconv.Atoi(st.Pop().(string))
	return val
}

func insertNVal(v1 interface{}, st *stack) {
	v2 := st.Peek()
	if IsNum(v1) && v2 != nil && IsNum(v2) {
		st.Pop()
		exp := st.Pop()
		nV := perfOp(v1, v2, exp)
		insertNVal(nV, st)
	} else {
		st.Push(v1)
	}

}

func perfOp(a, b, t interface{}) string {
	e := t.(string)
	switch e {
	case "+":
		return strconv.Itoa(num(a) + num(b))
	case "-":
		return strconv.Itoa(num(a) - num(b))
	case "*":
		return strconv.Itoa(num(a) * num(b))
	case "/":
		return strconv.Itoa(num(a) / num(b))
	default:
		return "0"
	}
}

func num(v interface{}) int {
	val, _ := strconv.Atoi(v.(string))
	return val

}
func IsNum(v interface{}) bool {

	if _, err := strconv.Atoi(v.(string)); err == nil {
		return true
	} else {
		return false
	}
}

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

func New() *stack {
	stk := new(stack)
	stk.lock = &sync.Mutex{}

	return stk
}

//func main() {
//	sl := []string{"5", "1", "2", "+", "4", "*", "+", "3", "-"}
//	fmt.Print(evalRPN(sl))
//}
