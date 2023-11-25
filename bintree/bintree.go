package bintree

// numericType is a generic type constrains consisting of Go numerics.
// NOTE: This is only for demonstration purpose only. There are a number
// of equivalent constraints in standard packages. One example is the
// `constraints` package. You can also use an alias to numeric constraints
// call `comparable`.
type numericType interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

// Node[N model.NumericType] of a binary tree
type Node[N numericType] interface {

	// Key returns the node's key
	Key() N

	// Count the number of duplicate keys
	Count() uint32

	// AddCount to number of duplicate keys
	AddCount(count uint32)

	// Left return the left child
	Left() Node[N]

	// SetLeft child
	SetLeft(left Node[N])

	// Right returns the right child
	Right() Node[N]

	// SetRight right child
	SetRight(left Node[N])
}

// NewNode[N model.NumericType] is a callback to an implementation to instantiate a node
type NewNode[N numericType] func(key N) Node[N]

// InsertNode[N model.NumericType] an operation to create a binary tree
func InsertNode[N numericType](newNode NewNode[N], root Node[N], key N) Node[N] {

	if root == nil {
		root = newNode(key)
	}

	if root.Key() == key {
		root.AddCount(root.Count() + 1)
	}

	if root.Key() < key {
		root.SetRight(InsertNode(newNode, root.Right(), key))
	}

	if root.Key() > key {
		root.SetLeft(InsertNode(newNode, root.Left(), key))
	}

	return root
}
