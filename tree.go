package ticktock

// OrderableFunc is a function that can be used to order the tree
// return value less than 0 means a < b
// return value equal to 0 means a == b
// return value greater than 0 means a > b
type OrderableFunc[T any] func(a, b T) int

// Node represents a node in the binary tree
type Node[T any] struct {
	value       T
	left, right *Node[T]
}

// NewNode creates a new node
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

// Insert inserts a node to the left/right node
// depending on the return value of the OrderableFunc
func (n *Node[T]) Insert(of OrderableFunc[T], value T) *Node[T] {
	if n == nil {
		return NewNode[T](value)
	}

	switch r := of(value, n.value); {
	case r < 0:
		n.left = n.left.Insert(of, value)
	case r > 0:
		n.right = n.right.Insert(of, value)
	}
	return n
}

// GetLeft returns the left child of the Node
func (n *Node[T]) GetLeft() *Node[T] {
	return n.left
}

// GetRight returns the right child of the node
func (n *Node[T]) GetRight() *Node[T] {
	return n.right
}

// GetValue returns the value stored in the node
func (n *Node[T]) GetValue() T {
	return n.value
}

// Find finds a node in the tree given an orderable
// function and a value
func (n *Node[T]) Find(of OrderableFunc[T], value T) *Node[T] {
	if n == nil {
		return nil
	}

	if of(value, n.value) == 0 {
		return n
	}

	if of(value, n.value) < 0 {
		return n.left.Find(of, value)
	}

	return n.right.Find(of, value)
}

// Tree has the root node and the OrderableFunc
type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

// CreateTree creates a new tree
func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

// GetRoot returns the root node of the tree
func (t *Tree[T]) GetRoot() *Node[T] {
	return t.root
}

// Find finds a node in the tree
func (t *Tree[T]) Find(value T) *Node[T] {
	return t.root.Find(t.f, value)
}

// Insert adds a child to the tree
func (t *Tree[T]) Insert(value T) {
	t.root = t.root.Insert(t.f, value)
}
