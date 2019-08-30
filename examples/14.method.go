package main

import "fmt"

type LinkedList struct {
	data int
	pre  *LinkedList
	next *LinkedList
}

type Index uint

func New(data int) *LinkedList {
	head := &LinkedList{
		data,
		nil,
		nil,
	}
	return head
}

func (l *LinkedList) Data() (int, bool) {
	if l == nil {
		return 0, false
	}
	return l.data, true
}

func (l *LinkedList) Previous() (*LinkedList, bool) {
	if l == nil {
		return nil, false
	}
	return l.pre, true
}

func (l *LinkedList) Next() (*LinkedList, bool) {
	if l == nil {
		return nil, false
	}
	return l.next, true
}

func (l *LinkedList) Head() *LinkedList {
	n := l
	for ; n.pre != nil; n = n.pre {
	}
	return n
}

func (l *LinkedList) Tail() *LinkedList {
	n := l
	for ; n.next != nil; n = n.next {
	}
	return n
}

func (l *LinkedList) Add(data int) *LinkedList {
	tail := l.Tail()
	node := &LinkedList{
		data,
		tail,
		nil,
	}
	tail.next = node
	return node
}

func (l *LinkedList) locate(index Index) *LinkedList {
	n := l.Head()
	for i := Index(0); n != nil && i < index; n, i = n.next, i+1 {
	}
	return n
}

func (l *LinkedList) Insert(data int, index Index) *LinkedList {
	node := &LinkedList{
		data,
		nil,
		nil,
	}
	n := l.locate(index)
	node.next = n
	node.pre = n.pre
	n.pre = node
	if node.pre != nil {
		node.pre.next = node
	}
	return node
}

func (l *LinkedList) Remove(index Index) *LinkedList {
	n := l.locate(index)
	if n != nil {
		if n.pre != nil {
			n.pre.next = n.next
		}
		if n.next != nil {
			n.next.pre = n.pre
		}
		n.pre = nil
		n.next = nil
	}
	return n
}

func (l *LinkedList) Update(data int, index Index) *LinkedList {
	n := l.locate(index)
	if n != nil {
		n.data = data
	}
	return n
}

func (l *LinkedList) PreOrder() {
	fmt.Println("PreOrder Traverse: ===")
	for n := l.Head(); n != nil; n = n.next {
		fmt.Println(n.data)
	}
}

func (l *LinkedList) PostOrder() {
	fmt.Println("PostOrder Traverse: ===")
	for n := l.Tail(); n != nil; n = n.pre {
		fmt.Println(n.data)
	}
}

func main() {
	l := New(10)
	for i := 20; i < 50; i += 10 {
		l.Add(i)
	}
	l.PreOrder()
	l.PostOrder()
	l.Insert(5, 0)
	l.Insert(25, 3)
	l.PreOrder()
	l.Remove(3)
	l.Remove(10)
	l.PostOrder()
	l.Update(100, 4)
	l.Update(1, 0)
	l.Update(500, 10)
	l.PreOrder()
}
