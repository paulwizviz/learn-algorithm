package linkedlist

// PNode use pointer to reference prior and next node
type PNode[V any] struct {
	Value V
	Next  *PNode[V]
	Prev  *PNode[V]
}

type Linkedlist[V any] interface {
	Append(V)
	Values() []V
}

type PSingleLinkedList[V any] struct {
	head *PNode[V]
}

func (p *PSingleLinkedList[V]) Append(value V) {
	n := &PNode[V]{
		Value: value,
	}
	if p.head == nil {
		p.head = n
		return
	}
	curr := p.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
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

func NewPSingleLinkedList[V any]() Linkedlist[V] {
	return &PSingleLinkedList[V]{
		head: nil,
	}
}
