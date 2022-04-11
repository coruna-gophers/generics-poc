package tree_test

import (
	"fmt"
	"testing"

	"github.com/coruna-gophers/generics-poc/tree"
	"github.com/stretchr/testify/assert"
)

func compareInt(a, b interface{}) int {
	if a == b {
		return 0
	}
	aInt, ok := a.(int)
	if !ok {
		panic(fmt.Errorf("'%s' is not an integer", a))
	}
	bInt, ok := b.(int)
	if !ok {
		panic(fmt.Errorf("'%s' is not an integer", b))
	}
	if aInt < bInt {
		return -1
	}
	if aInt == bInt {
		return 0
	}
	return 1
}

func TestInsert(t *testing.T) {
	tr := tree.New(compareInt)
	tr.Insert(1)
	tr.Insert(2)
	tr.Insert(3)
	tr.Insert(4)

	var result []interface{}
	tr.Walk(func(n *tree.Node) {
		result = append(result, n.Elem)
	})

	assert.Equal(t, []interface{}{1, 2, 3, 4}, result)
}
