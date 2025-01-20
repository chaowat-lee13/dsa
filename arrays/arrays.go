package arrays

import (
	"errors"
	"sort"
)

// building an array

type MyArray struct {
	length int
	data   []interface{}
}

func (ma *MyArray) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= ma.length {
		return nil, errors.New("index out of range")
	}
	return ma.data[idx], nil
}

func (ma *MyArray) Push(val interface{}) []interface{} {
	ma.data = append(ma.data, val)
	ma.length++
	return ma.data
}

func (ma *MyArray) Pop() (interface{}, error) {
	if len(ma.data) == 0 {
		return nil, errors.New("array is empty")
	}
	poppedItem := ma.data[ma.length-1]
	ma.data = ma.data[:ma.length-1]
	ma.length--
	return poppedItem, nil
}

func (ma *MyArray) Length() int {
	return ma.length
}

func ReverseString(s string) string {
	runes := []rune(s)
	if len(runes) < 2 {
		return s
	}

	placeHolder := make([]rune, len(runes))
	for i := 0; i < len(runes); i++ {
		placeHolder[i] = runes[len(runes)-1-i]
	}

	return string(placeHolder)
}

func MergeSortedArrays(a, b []int) []int {
	var mergedArray []int
	mergedArray = append(mergedArray, a...)
	mergedArray = append(mergedArray, b...)
	for i := 0; i < len(mergedArray)-1; i++ {
		swapped := false
		for j := 0; j < len(mergedArray)-i-1; j++ {
			if mergedArray[j] > mergedArray[j+1] {
				mergedArray[j], mergedArray[j+1] = mergedArray[j+1], mergedArray[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return mergedArray
}

func AnotherSolution(a, b []int) []int {
	var sortedArray []int
	sortedArray = append(sortedArray, a...)
	sortedArray = append(sortedArray, b...)

	sort.Ints(sortedArray)

	return sortedArray
}
