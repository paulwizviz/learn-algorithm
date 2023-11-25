package bintree

// default implementation of a binary node using pointers.
// NOTE: This is not serializable.
type defaultNode[N numericType] struct {
	key   N
	count uint32
	left  Node[N]
	right Node[N]
}

func (p *defaultNode[N]) Key() N {
	return p.key
}

func (d *defaultNode[N]) Count() uint32 {
	return d.count
}

func (d *defaultNode[N]) AddCount(count uint32) {
	d.count = count
}

func (p *defaultNode[N]) Left() Node[N] {
	return p.left
}

func (p *defaultNode[N]) SetLeft(left Node[N]) {
	p.left = left
}

func (p *defaultNode[N]) Right() Node[N] {
	return p.right
}

func (p *defaultNode[N]) SetRight(right Node[N]) {
	p.right = right
}

// NewDefaultNode instantiate a reference to adefault implementation
// of a Node
func NewDefaultNode[N numericType](key N) Node[N] {
	return &defaultNode[N]{
		key: key,
	}
}
