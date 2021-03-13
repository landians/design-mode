package iterator

import "errors"

// FIFO 队列接口
type ILinkedList interface {
	Size() int
	Push(it interface{})
	Poll() (error, interface{})
	Iterator() IIterator
}

// 迭代器接口，More() 用于判断是否有更多的元素，Next() 用于获取下一个元素
type IIterator interface {
	More() bool
	Next() (error, interface{})
}

// 实现 ILinkedList 接口
type xLinkedList struct {
	size int
	head *xLinkedNode
	tail *xLinkedNode
}

func newLinkedList() ILinkedList {
	return &xLinkedList{}
}

// linkedList 的节点
type xLinkedNode struct {
	value interface{}
	next  *xLinkedNode
}

func newLinkedNode(v interface{}) *xLinkedNode {
	return &xLinkedNode{
		value: v,
	}
}

func (l *xLinkedList) Size() int {
	return l.size
}

func (l *xLinkedList) Push(v interface{}) {
	node := newLinkedNode(v)
	if l.head == nil {
		l.head, l.tail = node, node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

func (l *xLinkedList) Poll() (error, interface{}) {
	if l.size <= 0 {
		return errors.New("empty list"), nil
	}

	node := l.head
	l.head = l.head.next
	l.size--
	return nil, node.value
}

func (l *xLinkedList) Iterator() IIterator {
	return newLinkedListIterator(l)
}

// 队列迭代器，实现 IIterator 接口
type xLinkedListIterator struct {
	list    *xLinkedList
	current *xLinkedNode
}

func newLinkedListIterator(list *xLinkedList) IIterator {
	return &xLinkedListIterator{
		list:    list,
		current: list.head,
	}
}

func (l *xLinkedListIterator) More() bool {
	return l.current != nil
}

func (l *xLinkedListIterator) Next() (error, interface{}) {
	node := l.current
	if node == nil {
		return errors.New("no more elements"), nil
	}

	l.current = l.current.next
	return nil, node.value
}

