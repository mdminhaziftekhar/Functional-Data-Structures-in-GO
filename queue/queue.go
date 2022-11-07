package main

import (
	"fmt"
)

type Queue struct {
	front       int
	actualQueue []int
}

func (s Queue) IsEmpty() bool {
	return s.front < 0
}

func noOfInToBeEnqueued(vals ...int) int {
	emptySlice := []int{}
	s := append(emptySlice, vals...)
	n := len(s)
	return n
}

func (q Queue) enqueue(vals ...int) (Queue, []int) {
	noOfElements := noOfInToBeEnqueued(vals...)
	q.front = q.front + noOfElements
	q.actualQueue = append(q.actualQueue, vals...)
	return q, q.actualQueue
}

func (q Queue) Dequeue() (Queue, []int) {
	element := q.actualQueue[0]
	fmt.Println("Dequeued:", element)
	q.front--
	return q, q.actualQueue[1:]
}

func (q Queue) Front() int {
	if !q.IsEmpty() {
		val := q.actualQueue[q.front]
		return val
	}
	return 0
}

func main() {
	q := Queue{front: -1}
	newQueue, actualQueue := q.enqueue(8, 10)
	fmt.Println(actualQueue)

	newQueue2, actualQueue2 := newQueue.enqueue(90)
	fmt.Println(newQueue2, actualQueue2)

	fmt.Println("After dequeue 8: ", newQueue2.actualQueue)
	afterDequeue, dequedElement := newQueue2.Dequeue()
	fmt.Println("After dequeue: ", afterDequeue, dequedElement)

	fmt.Println("Front of the queue: ", afterDequeue.Front())
}
