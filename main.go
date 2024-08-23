package main

import (
	"dsa/arrays"
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

	stringTest := "I AM Chaowat World"
	reverseString := arrays.ReverseString(stringTest)
	fmt.Println(reverseString)

	a := []int{1, 2, 11, 31, 5, 9, 10}
	b := []int{23, 6, 7, 8, 99}
	c := arrays.MergeSortedArrays(a, b)

	fmt.Println(c)

}
