package main

import "fmt"

type Stack struct {
	items []int
}

type Queue struct {
	items []int
}

//push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

//enqueue

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//deque
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println("Push and Pop in stack")
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
	myQueue := Queue{}
	fmt.Println("Enqueue and Dequeue in queue")
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
