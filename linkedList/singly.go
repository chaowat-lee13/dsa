package linkedList

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (ll *LinkedList) Init(data int) {
	newNode := &Node{data: data}
	ll.head = newNode
	ll.tail = newNode
	ll.length = 1
}

func (ll *LinkedList) Prepend(data int) {
	newNode := &Node{data: data}
	newNode.next = ll.head
	ll.head = newNode
	if ll.tail == nil {
		ll.tail = newNode
	}
	ll.length++
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data}
	ll.tail.next = newNode
	ll.tail = newNode
	ll.length++
}

func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (ll *LinkedList) Insert(index, data int) {
	if index < 0 {
		return
	}

	newNode := &Node{data: data}

	if index == 0 {
		ll.Prepend(data)
		return
	}
	if index >= ll.length {
		ll.Append(data)
		return
	}

	// {head: 2, next: {head: 4, next: {head: 8, next: {head: 9: next: nil}}}}
	current := ll.head
	currentIndex := 0
	for current != nil && currentIndex < index-1 {
		current = current.next
		currentIndex++
	}

	newNode.next = current.next
	current.next = newNode
	ll.length++
}

func (ll *LinkedList) Remove(index int) {
	if index < 0 || index >= ll.length {
		return
	}

	if index == 0 {
		ll.head = ll.head.next
		ll.length--
		if ll.length == 0 {
			ll.tail = nil
		}
		return
	}

	current := ll.head
	currentIndex := 0
	for current.next != nil && currentIndex < index-1 {
		current = current.next
		currentIndex++
	}

	// {head: 8, next: {head: 9, next: {head: 11: next: {head: 12, next: nil}}}}
	if current.next == nil {
		return
	}

	//{head: 11: next: {head: 12, next: nil}}
	current.next = current.next.next
	ll.length--

	if current.next == nil {
		ll.tail = current
	}
}

func (ll *LinkedList) Reverse() {
	var prev *Node = nil
	current := ll.head
	next := &Node{}

	for current != nil {
		next = current.next // Store the next node
		current.next = prev // Reverse the current node's pointer
		prev = current      // Move 'prev' one step forward
		current = next      // Move 'current' one step forward
	}

	ll.head = prev // Update the head to point to the new head (last node)

	ll.PrintList()
}
