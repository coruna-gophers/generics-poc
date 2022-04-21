package tree

// -1: a smaller than b
// 0: a equals to b
// 1: a greater than b
type Comparator func(keyA, keyB interface{}) int

type Tree struct {
	root    *node
	compare Comparator
}

type node struct {
	right *node
	left  *node
	entry *entry
}

type entry struct {
	key   interface{}
	value interface{}
}

func New(comparator Comparator) *Tree {
	return &Tree{
		root:    &node{},
		compare: comparator,
	}
}

func (t *Tree) Insert(key, value interface{}) {
	entry := &entry{
		key:   key,
		value: value,
	}
	if t.root.entry == nil {
		t.root.entry = entry
		return
	}
	t.insert(t.root, entry)
}

func (t *Tree) insert(n *node, entry *entry) {
	c := t.compare(entry.key, n.entry.key)
	if c < 0 || c == 0 {
		if n.left == nil {
			n.left = &node{
				entry: entry,
			}
		} else {
			t.insert(n.left, entry)
		}
	}
	if c > 0 {
		if n.right == nil {
			n.right = &node{
				entry: entry,
			}
		} else {
			t.insert(n.right, entry)
		}
	}
}

func (t *Tree) Find(key interface{}) (interface{}, bool) {
	node, found := t.find(t.root, key)
	if found {
		return node.entry.value, true
	}
	return nil, false
}

func (t *Tree) find(n *node, key interface{}) (*node, bool) {
	c := t.compare(key, n.entry.key)
	if c == 0 {
		return n, true
	}
	if c < 0 && n.left != nil {
		return t.find(n.left, key)
	}
	if c > 0 && n.right != nil {
		return t.find(n.right, key)
	}
	return nil, false
}

type Walker func(key, value interface{})

// Walk traverses the tree in order
func (t *Tree) Walk(cb Walker) {
	t.walk(t.root, cb)
}

func (t *Tree) walk(n *node, cb Walker) {
	if n.left != nil {
		cb(n.entry.key, n.entry.value)
		t.walk(n.left, cb)
		return
	}
	if n.left == nil && n.right == nil {
		cb(n.entry.key, n.entry.value)
		return
	}
	if n.right != nil {
		cb(n.entry.key, n.entry.value)
		t.walk(n.right, cb)
		return
	}
}
