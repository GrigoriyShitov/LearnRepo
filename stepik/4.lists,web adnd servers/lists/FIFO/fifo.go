package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) { //push to the end of the list
	//
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} { //remove from front of list
	// ...
	ret := queue.Front()
	if ret != nil {
		queue.Remove(ret)
		return ret.Value
	} else {
		return nil
	}

}

func printQueue(queue *list.List) { //print the list
	// ...
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
}

func main() {

	queue := list.New()

	Push(1, queue)

	Push(2, queue)

	Push(3, queue)

	printQueue(queue) // 123

	Pop(queue)

	printQueue(queue) // в ту же строку: 23

}
