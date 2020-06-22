package main

import (
	"container/list"
	"fmt"
)

func main() {
	tmpList := list.New()

	for i := 1; i <= 10; i++ {
		tmpList.PushBack(i)
		fmt.Println(tmpList)
	}

	first := tmpList.PushFront(1)
	fmt.Println(first)
	fmt.Println(tmpList)
	tmpList.Remove(first)
	fmt.Println(tmpList)
}
