package pq

import (
	"github.com/Rongcong/algorithms/data-structures/heap"
	"github.com/Rongcong/algorithms/data-structures/queue"
)

// Item ...
type Item struct {
	Value    interface{}
	Priority int
}

// NewItem ...
func NewItem(value interface{}, priority int) (i *Item) {
	return &Item{
		Value:    value,
		Priority: priority,
	}
}

// Less ...
func (x Item) Less(than heap.Item) bool {
	return x.Priority < than.(Item).Priority
}

// PQ ...
type PQ struct {
	data heap.Heap
}

// NewMax ...
func NewMax() (q *PQ) {
	return &PQ{
		data: *heap.NewMax(),
	}
}

// NewMin ...
func NewMin() (q *PQ) {
	return &PQ{
		data: *heap.NewMin(),
	}
}

// Len ...
func (pq *PQ) Len() int {
	return pq.data.Len()
}

// Insert ...
func (pq *PQ) Insert(el Item) {
	pq.data.Insert(heap.Item(el))
}

// Extract ...
func (pq *PQ) Extract() (el Item) {
	return pq.data.Extract().(Item)
}

// ChangePriority ...
func (pq *PQ) ChangePriority(val interface{}, priority int) {
	var storage = queue.New()

	popped := pq.Extract()

	for val != popped.Value {
		if pq.Len() == 0 {
			panic("Item not found")
		}

		storage.Push(popped)
		popped = pq.Extract()
	}

	popped.Priority = priority
	pq.data.Insert(popped)

	for storage.Len() > 0 {
		pq.data.Insert(storage.Shift().(heap.Item))
	}
}
