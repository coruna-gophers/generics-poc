package gtree

// -1: a smaller than b
// 0: a equals to b
// 1: a greater than b

type Comparator[K any] func(keyA K, keyB K) int

type Tree[K, V any] struct {
	root    *node[K, V]
	compare Comparator[K]
}

type node[K, V any] struct {
	right *node[K, V]
	left  *node[K, V]
	entry *entry[K, V]
}

type entry[K, V any] struct {
	key   K
	value V
}

func New[K, V any](comparator Comparator[K]) *Tree[K, V] {
	return &Tree[K, V]{
		root:    &node[K, V]{},
		compare: comparator,
	}
}

func (t *Tree[K, V]) Insert(key K, value V) {
	entry := &entry[K, V]{
		key:   key,
		value: value,
	}
	if t.root.entry == nil {
		t.root.entry = entry
		return
	}
	t.insert(t.root, entry)
}

func (t *Tree[K, V]) insert(n *node[K, V], entry *entry[K, V]) {
	c := t.compare(entry.key, n.entry.key)
	if c < 0 || c == 0 {
		if n.left == nil {
			n.left = &node[K, V]{
				entry: entry,
			}
		} else {
			t.insert(n.left, entry)
		}
	}
	if c > 0 {
		if n.right == nil {
			n.right = &node[K, V]{
				entry: entry,
			}
		} else {
			t.insert(n.right, entry)
		}
	}
}

func (t *Tree[K, V]) Find(key K) (result V, found bool) {
	node, found := t.find(t.root, key)
	if found {
		return node.entry.value, true
	}
	return
}

func (t *Tree[K, V]) find(n *node[K, V], key K) (*node[K, V], bool) {
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

type Walker[K, V any] func(key K, value V)

// Walk traverses the tree in order
func (t *Tree[K, V]) Walk(cb Walker[K, V]) {
	t.walk(t.root, cb)
}

func (t *Tree[K, V]) walk(n *node[K, V], cb Walker[K, V]) {
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
