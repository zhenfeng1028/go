package main

import (
	"container/list"
	"fmt"
)

type HashTable struct {
	capacity int
	table    []*list.List
}

type KVP struct {
	key   int
	value int
}

func checkPrime(n int) bool {
	// 0和1都不是素数
	if n == 0 || n == 1 {
		return false
	}
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getPrime(n int) int {
	if n%2 == 0 {
		n++
	}
	for !checkPrime(n) {
		n += 2
	}
	return n
}

func NewHashTable(size int) *HashTable {
	newSize := getPrime(size)
	ht := &HashTable{}
	ht.capacity = newSize
	ht.table = make([]*list.List, newSize)
	for i := 0; i < ht.capacity; i++ {
		ht.table[i] = list.New()
	}
	return ht
}

func (t *HashTable) hashFunc(key int) int {
	return key % t.capacity
}

func (t *HashTable) insertItem(item KVP) {
	index := t.hashFunc(item.key)
	t.table[index].PushBack(item)
}

func (t *HashTable) deleteItem(item KVP) {
	index := t.hashFunc(item.key)
	l := t.table[index]
	var e *list.Element
	for e = l.Front(); e != nil; e = e.Next() {
		kvp, _ := e.Value.(KVP)
		if kvp.key == item.key && kvp.value == item.value {
			break
		}
	}
	t.table[index].Remove(e)
}

func (t *HashTable) displayHash() {
	for i := 0; i < t.capacity; i++ {
		fmt.Printf("table[%d]", i)
		l := t.table[i]
		for e := l.Front(); e != nil; e = e.Next() {
			kvp, _ := e.Value.(KVP)
			fmt.Print(" -> ", kvp.value)
		}
		fmt.Println()
	}
}

func main() {
	kvps := []KVP{
		{231, 123},
		{321, 432},
		{212, 523},
		{321, 43},
		{433, 423},
		{262, 111},
	}

	size := len(kvps)

	ht := NewHashTable(size)

	for i := 0; i < size; i++ {
		ht.insertItem(kvps[i])
	}

	ht.deleteItem(KVP{321, 432})
	ht.displayHash()
}
