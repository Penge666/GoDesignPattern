package main

import "fmt"

// ----------- 自定义容器 -------------
type IntCollection struct {
	items []int
}

func NewIntCollection(items []int) *IntCollection {
	return &IntCollection{items: items}
}

func (c *IntCollection) Iterator() *IntIterator {
	return &IntIterator{
		collection: c,
		index:      0,
	}
}

// ----------- 迭代器结构体 -------------
type IntIterator struct {
	collection *IntCollection
	index      int
}

func (it *IntIterator) HasNext() bool {
	return it.index < len(it.collection.items)
}

func (it *IntIterator) Next() int {
	val := it.collection.items[it.index]
	it.index++
	return val
}

// ----------- 使用示例 -------------
func main() {
	col := NewIntCollection([]int{10, 20, 30, 40})
	it := col.Iterator()

	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
