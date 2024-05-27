package linkedlist

// PNode use pointer to reference prior and next node
type PNode[V any] struct {
	ID    uint
	Value V
	Next  *PNode[V]
	Prev  *PNode[V]
}

type PLinkedlist[V any] interface {
	Append(V) // Append new node to list
	Search(uint) *PNode[V]
	Values() []V // Show all values in the list
}

type PSingleLinkedList[V any] struct {
	id   uint
	head *PNode[V]
}

func (p *PSingleLinkedList[V]) Append(value V) {
	n := &PNode[V]{
		ID:    p.id,
		Value: value,
	}
	if p.head == nil {
		p.head = n
		p.id = p.id + 1
		return
	}
	curr := p.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
	p.id = p.id + 1
}

func (p PSingleLinkedList[V]) Search(id uint) *PNode[V] {
	curr := p.head
	if curr == nil {
		return nil
	}
	for curr.Next != nil {
		if curr.ID == id {
			return curr
		}
		curr = curr.Next
	}
	if curr.ID == id {
		return curr
	}
	return nil
}

func (p *PSingleLinkedList[V]) Values() []V {
	values := []V{}
	curr := p.head
	for curr.Next != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}
	values = append(values, curr.Value)
	return values
}

func NewPSingleLinkedList[V any]() PLinkedlist[V] {
	return &PSingleLinkedList[V]{
		head: nil,
	}
}
