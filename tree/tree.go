package tree

// -1: a smaller than b
// 0: a equals to b
// 1: a greater than b
type Comparator func(a, b interface{}) int

type Tree struct {
	Root    *Node
	Compare Comparator
}

type Node struct {
	Right *Node
	Left  *Node
	Elem  interface{}
}

func New(comparator Comparator) *Tree {
	return &Tree{
		Root:    &Node{},
		Compare: comparator,
	}
}

func (t *Tree) Insert(elem interface{}) {
	if t.Root.Elem == nil {
		t.Root.Elem = elem
		return
	}

	t.insert(t.Root, elem)
}

func (t *Tree) insert(node *Node, elem interface{}) {
	c := t.Compare(elem, node.Elem)
	if c < 0 || c == 0 {
		if node.Left == nil {
			left := &Node{
				Elem: elem,
			}
			node.Left = left
		} else {
			t.insert(node.Left, elem)
		}
	}
	if c > 0 {
		if node.Right == nil {
			right := &Node{
				Elem: elem,
			}
			node.Right = right
		} else {
			t.insert(node.Right, elem)
		}
	}
}

func (t *Tree) Delete(node *Node, elem interface{}) {

}

func (t *Tree) Find(node *Node, elem interface{}) {

}

// Walk traverses the tree in order
func (t *Tree) Walk(cb func(n *Node)) {
	t.walk(t.Root, cb)
}

func (t *Tree) walk(node *Node, cb func(n *Node)) {
	if node.Left != nil {
		cb(node)
		t.walk(node.Left, cb)
		return
	}
	if node.Left == nil && node.Right == nil {
		cb(node)
		return
	}
	if node.Right != nil {
		cb(node)
		t.walk(node.Right, cb)
		return
	}
}
