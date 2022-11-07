package main

import "fmt"

type stack struct {
	top         int
	actualStack []int
}

func (s stack) IsEmpty() bool {
	return s.top < 0
}

func noOfInToBePushed(vals ...int) int {
	emptySlice := []int{}
	s := append(emptySlice, vals...)
	n := len(s)
	return n
}

func (s stack) Push(vals ...int) (stack, []int) {
	noOfArgs := noOfInToBePushed(vals...)
	s.top = s.top + noOfArgs
	s.actualStack = append(s.actualStack, vals...)
	return s, s.actualStack
}

func (s stack) Pop() (stack, []int) {
	if s.top < 0 {
		return s, nil
	}


	s2 := (s.actualStack)[:s.top]

	s3 := stack{top: -1}

	s3.top = len(s2) - 1
	s3.actualStack = s2


	return s3, s3.actualStack
}

func (s stack) Top() int {
	if !s.IsEmpty() {
		val := s.actualStack[s.top]
		return val
	}
	return 0
}

func main() {
	s := stack{top: -1}
	stk, actaulStk := s.Push(1, 3, 4, 6)
	fmt.Println(actaulStk)

	//x := 9
	stkNewPush, actualStkAfterPush := stk.Push(100, 200)
	fmt.Println("New stack object after push: ", stkNewPush)
	fmt.Println("Stack after pushing 9: ", actualStkAfterPush)

	stkAfterPop, actualStkAfterPop := stkNewPush.Pop()
	fmt.Println("Stack after Pop:", stkAfterPop)
	fmt.Println("After pop: ", actualStkAfterPop)

	topVal := stkAfterPop.Top()
	fmt.Println("Top value: ", topVal)
}
