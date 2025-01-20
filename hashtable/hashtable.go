package hashtable

import (
	"errors"
	"strconv"
)

// Define the structure of a key-value pair
type keyValuePair struct {
	key   string
	value int
}

// Define the hash table structure
type HashTable struct {
	table    [][]keyValuePair
	size     int
	capacity int
}

// Create a new hash table
func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		table:    make([][]keyValuePair, capacity),
		size:     0,
		capacity: capacity,
	}
}

// A simple hash function (you can use more sophisticated ones)
func (ht *HashTable) hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ht.capacity
}

// Insert a key-value pair into the hash table
func (ht *HashTable) Insert(key string, value int) {
	index := ht.hash(key)
	ht.table[index] = append(ht.table[index], keyValuePair{key, value})
	ht.size++
}

// Get the value associated with a key
func (ht *HashTable) Get(key string) int {
	index := ht.hash(key)
	for _, kv := range ht.table[index] {
		if kv.key == key {
			return kv.value
		}
	}
	return 0
}

// Remove a key-value pair from the hash table
func (ht *HashTable) Remove(key string) bool {
	index := ht.hash(key)
	for i, kv := range ht.table[index] {
		if kv.key == key {
			ht.table[index] = append(ht.table[index][:i], ht.table[index][i+1:]...)
			ht.size--
			return true
		}
	}
	return false
}

func FirstRecurring(array []int) (int, error) {
	recurring := make(map[string]int)

	for _, value := range array {
		recurring[strconv.Itoa(value)] += 1
		if recurring[strconv.Itoa(value)] == 2 {
			return value, nil
		}
	}

	return 0, errors.New("undefined")
}
