package main

import "fmt"

// LinkedList 双链表结构体定义
type LinkedList struct {
	Data int
	Pre  *LinkedList // 上一个节点
	Next *LinkedList // 下一个结点
}

// Head 初始化双链表头，是一个链表结点指针
var Head = &LinkedList{
	10,
	nil,
	nil,
}

// Tail 链表尾，初始化为链表头
var Tail = Head

// Add 添加新的元素到链表尾部
func Add(data int) {
	// 新节点，取址
	node := &LinkedList{
		data,
		Tail,
		nil,
	}
	// 接入链表
	Tail.Next = node
	// 链表尾设置为新节点
	Tail = node
}

// IncList 将链表中每个元素增加5
func IncList() {
	for n := Head; n != nil; n = n.Next {
		n.Data += 5
	}
}

func main() {
	// 向链表中添加三个元素 [20, 30, 40]
	for i := 20; i < 50; i += 10 {
		Add(i)
	}
	// 前序遍历
	fmt.Println("PreOrder Traverse: ===")
	// 标准遍历循环
	for n := Head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}
	IncList()
	// 后序遍历
	fmt.Println("PostOrder Traverse: ===")
	for n := Tail; n != nil; n = n.Pre {
		fmt.Println(n.Data)
	}
}
