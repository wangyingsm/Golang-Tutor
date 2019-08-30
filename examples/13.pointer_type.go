package main

import "fmt"

type LinkedList struct {
	Data int
	Pre  *LinkedList
	Next *LinkedList
}

var Head = &LinkedList{
	10,
	nil,
	nil,
}

var Tail = Head

func Add(data int) {
	node := &LinkedList{
		data,
		Tail,
		nil,
	}
	Tail.Next = node
	Tail = node
}

func IncList() {
	for n := Head; n != nil; n = n.Next {
		n.Data += 5
	}
}

func main() {
	for i := 20; i < 50; i += 10 {
		Add(i)
	}
	fmt.Println("PreOrder Traverse: ===")
	for n := Head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}
	IncList()
	fmt.Println("PostOrder Traverse: ===")
	for n := Tail; n != nil; n = n.Pre {
		fmt.Println(n.Data)
	}
}
