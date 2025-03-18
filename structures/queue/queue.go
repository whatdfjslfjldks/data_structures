package main

import "fmt"

// 用两个栈模拟队列

// Stack 定义栈结构
type Stack struct {
	elements []int
}

// Push 向栈中推入一个元素
func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

// Pop 从栈中弹出一个元素
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Queue 用两个栈模拟队列
type Queue struct {
	inStack  Stack // 用来入队
	outStack Stack // 用来出队
}

// Enqueue 入队
func (q *Queue) Enqueue(element int) {
	q.inStack.Push(element)
}

// Dequeue 出队 如果outStack不为空，则直接从outStack中弹出一个元素
// 如果outStack为空，则从inStack中弹出所有元素，然后依次入outStack
func (q *Queue) Dequeue() int {
	if q.outStack.IsEmpty() {
		for !q.inStack.IsEmpty() {
			q.outStack.Push(q.inStack.Pop())
		}
	}
	return q.outStack.Pop()
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	a := queue.Dequeue()
	fmt.Println(a)
}
