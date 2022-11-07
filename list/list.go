package main

import "fmt"

type List struct {
	Data interface{}
}

func (List List) GetData() interface{} {
	return List.Data
}

func (List *List) SetData(data interface{}) {
	List.Data = data
}

func cons(list *List, toAdd int) []int {
	var data []int = list.GetData().([]int)
	list2 := append(data, toAdd)
	return list2
}

func car(list []int) int {
	head := list[0]
	return head
}

func cdr(list []int) []int {
	tail := list[1:]
	return tail
}

func main() {
	list := &List{}
	list.SetData([]int{1, 2, 3, 4})

	list2 := cons(list, 10)

	fmt.Println(list2)
	fmt.Println(car(list2))
	fmt.Println(cdr(list2))

}
