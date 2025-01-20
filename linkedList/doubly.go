package linkedList

import "fmt"

type DoublyNode struct {
	data int
	next *DoublyNode
	prev *DoublyNode
}

type DoublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int
}

func (dl *DoublyLinkedList) Init(data int) {
	newNode := &DoublyNode{data: data}
	dl.head = newNode
	dl.tail = newNode
	dl.length = 1
}

func (dl *DoublyLinkedList) Append(data int) {
	newNode := &DoublyNode{data: data}
	newNode.prev = dl.tail
	dl.tail.next = newNode
	dl.tail = newNode
	dl.length++
}

func (dl *DoublyLinkedList) Prepend(data int) {
	newNode := &DoublyNode{data: data}
	newNode.next = dl.head
	dl.head.prev = newNode
	dl.head = newNode
	if dl.tail == nil {
		dl.tail = newNode
	}
	dl.length++
}

func (dl *DoublyLinkedList) PrintList() {
	current := dl.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (dl *DoublyLinkedList) Insert(index, data int) {
	if index < 0 {
		return
	}

	newNode := &DoublyNode{data: data}

	if index == 0 {
		dl.Prepend(data)
		return
	}
	if index >= dl.length {
		dl.Append(data)
		return
	}

	current := dl.head
	currentIndex := 0
	for current != nil && currentIndex < index-1 {
		current = current.next
		currentIndex++
	}

	newNode.next = current.next
	newNode.prev = current
	current.next.prev = newNode
	current.next = newNode
	dl.length++
}

func (dl *DoublyLinkedList) Remove(index int) {
	if index < 0 || index > dl.length {
		return
	}

	if index == 0 {
		dl.head = dl.head.next
		if dl.head != nil {
			dl.head.prev = nil
		}
		dl.length--
		if dl.length == 0 {
			dl.tail = nil
		}
		return
	}

	current := dl.head
	currentIndex := 0
	for current.next != nil && currentIndex < index-1 {
		current = current.next
		currentIndex++
	}

	if current.next == nil {
		return
	}

	current.next = current.next.next
	if current.next != nil {
		current.next.prev = current
	} else {
		dl.tail = current
	}
	dl.length--
}

func (dl *DoublyLinkedList) PrintAll() {
	current := dl.head
	for current != nil {
		fmt.Printf("Node: %p, Data: %d, Prev: %p, Next: %p\n", current, current.data, current.prev, current.next)
		current = current.next
	}
	fmt.Println()
}

func (dl *DoublyLinkedList) Reverse() {
	temp := &DoublyNode{} // Temporary node for swapping
	current := dl.head

	for current != nil {
		// Swap prev and next pointers
		temp = current.prev
		current.prev = current.next
		current.next = temp

		current = current.prev // Move to the next node (which was previously the previous)
	}

	// Swap head and tail
	temp = dl.head
	dl.head = dl.tail
	dl.tail = temp

	dl.PrintList()
}
