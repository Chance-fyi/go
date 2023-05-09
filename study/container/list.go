package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println(l)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	fmt.Println(l)
	fmt.Println(l.Front())
	fmt.Println(l.Front().Next())
	fmt.Println(l.Front().Next().Next())
	fmt.Println(l.Front().Next().Next().Next())
	fmt.Println(l.Front().Next().Next().Next().Next())
	fmt.Println(l.Front().Next().Next().Next().Next().Next())
	fmt.Println(l.Front().Prev())
}
