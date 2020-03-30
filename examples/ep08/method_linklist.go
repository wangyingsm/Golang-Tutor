package main

import "fmt"

// LinkedList 双链表结构体
type LinkedList struct {
	data int
	pre  *LinkedList
	next *LinkedList
}

// Index 链表序号类型
type Index uint

// New 构造函数，返回一个双链表头指针
func New(data int) *LinkedList {
	head := &LinkedList{
		data,
		nil,
		nil,
	}
	return head
}

// Data 获取当前链表指针数据的Getter
func (l *LinkedList) Data() (int, bool) {
	if l == nil {
		return 0, false
	}
	return l.data, true
}

// Previous 链表向前移动一个元素
func (l *LinkedList) Previous() (*LinkedList, bool) {
	if l == nil {
		return nil, false
	}
	return l.pre, true
}

// Next 链表向后移动一个元素
func (l *LinkedList) Next() (*LinkedList, bool) {
	if l == nil {
		return nil, false
	}
	return l.next, true
}

// Head 回到链表头
func (l *LinkedList) Head() *LinkedList {
	n := l
	for ; n.pre != nil; n = n.pre {
	}
	return n
}

// Tail 到链表尾
func (l *LinkedList) Tail() *LinkedList {
	n := l
	for ; n.next != nil; n = n.next {
	}
	return n
}

// Add 在链表结尾添加元素
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

// locate 定位到链表的第index个元素
func (l *LinkedList) locate(index Index) *LinkedList {
	n := l.Head()
	for i := Index(0); n != nil && i < index; n, i = n.next, i+1 {
	}
	return n
}

// Insert 在链表第index个元素前面插入一个新元素
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

// Remove 删除链表第index个元素
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

// Update 更新链表第index个元素
func (l *LinkedList) Update(data int, index Index) *LinkedList {
	n := l.locate(index)
	if n != nil {
		n.data = data
	}
	return n
}

// PreOrder 链表前序遍历
func (l *LinkedList) PreOrder() {
	fmt.Println("PreOrder Traverse: ===")
	for n := l.Head(); n != nil; n = n.next {
		fmt.Println(n.data)
	}
}

// PostOrder 链表后序遍历
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
