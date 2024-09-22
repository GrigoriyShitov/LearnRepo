package main

import (
	"container/list"
	"fmt"
)

func ReverseList(l *list.List) *list.List {
	// Здесь ваш код
	reversedList := list.New()
	for temp := l.Front(); temp != nil; temp = temp.Next() {
		reversedList.PushFront(temp.Value)
	}
	return reversedList
}

func printList(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
	fmt.Println()
}

func main() {

	task := list.New()

	for i := 0; i < 10; i++ {
		task.PushBack(i)
	}
	printList(task)
	// 0123456789

	printList(ReverseList(task))
	// 9876543210

}
