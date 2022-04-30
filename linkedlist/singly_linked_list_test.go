package linkedlist

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtHead(7)
	list.DeleteAtIndex(1)
	list.Traverse()
}

type MyLinkedList struct {
	head *Node
	len  int
}

type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, len: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if this.len == 0 || index >= this.len {
		return -1
	}
	node := this.GetNode(index)
	return node.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)

}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.len, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len || index < 0 {
		return
	}

	if this.len == 0 {
		node := &Node{val: val}
		this.head = node
		this.len++
		return
	}

	if index == 0 {
		node := &Node{val: val, next: this.head}
		this.head = node
		this.len++
		return
	}

	if index == this.len {
		last := this.GetNode(this.len - 1)
		node := &Node{val: val}
		last.next = node
		this.len++
		return
	}

	pre := this.GetNode(index - 1)
	node := &Node{val: val, next: pre.next}
	pre.next = node
	this.len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}

	if index == 0 {
		this.head = this.head.next
		this.len--
		return
	}

	pre := this.GetNode(index - 1)
	pre.next = pre.next.next
	this.len--
}

func (this *MyLinkedList) GetNode(index int) *Node {
	node := this.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (this *MyLinkedList) Traverse() {
	cur := this.head
	for cur.next != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
	fmt.Println(cur.val)
}
