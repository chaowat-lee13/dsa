package main

import (
	"dsa/linkedList"
	"fmt"
)

func main() {
	//myArr := arrays.MyArray{}
	//
	//myArr.Push("A")
	//
	//myArr.Push("B")
	//
	//val, err := myArr.Get(1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(val)
	//
	//if err = myArr.Pop(); err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(myArr)

	//stringTest := "I AM Chaowat World"
	//reverseString := arrays.ReverseString(stringTest)
	//fmt.Println(reverseString)
	//
	//a := []int{1, 2, 11, 31, 5, 9, 10}
	//b := []int{23, 6, 7, 8, 99}
	//c := arrays.MergeSortedArrays(a, b)
	//d := arrays.AnotherSolution(a, b)
	//
	//fmt.Println(c)
	//
	//fmt.Println(d)

	// Making My Hash Table
	//ht := hashtable.NewHashTable(10)
	//ht.Insert("apple", 100)
	//ht.Insert("banana", 223)
	//ht.Insert("cherry", 35000)
	//
	//fmt.Println(ht.Get("banana"))
	//fmt.Println(ht.Get("grape"))
	//
	//ht.Remove("banana")
	//fmt.Println(ht.Get("banana"))
	//fmt.Println(ht.Get("cherry"))
	// Hash Table Exercise
	//a := []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	//resultA, err := hashtable.FirstRecurring(a)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("The result of recurring function A: ", resultA)
	//
	//b := []int{7, 5, 1, 2, 3, 5, 1, 2, 4}
	//resultB, err := hashtable.FirstRecurring(b)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("The result of recurring function B: ", resultB)
	//
	//c := []int{2, 3, 4, 5}
	//resultC, err := hashtable.FirstRecurring(c)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("The result of recurring function C: ", resultC)
	//}

	// Singly Linked List
	ll := linkedList.LinkedList{}
	ll.Init(4)
	ll.Append(8)
	ll.Append(9)
	ll.Append(12)
	ll.Insert(2, 11)

	fmt.Println("Singly Linked list:")
	ll.PrintList()

	ll.Reverse()

	// Doubly Linked List

	dl := linkedList.DoublyLinkedList{}
	dl.Init(4)
	dl.Append(8)
	dl.Append(9)
	dl.Append(11)
	dl.Append(12)
	dl.Prepend(1)
	dl.Remove(2)
	dl.Insert(3, 88)
	dl.Remove(5)

	fmt.Println("Doubly Linked list:")
	dl.PrintList()
	dl.Reverse()
}
