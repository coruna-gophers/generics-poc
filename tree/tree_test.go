package tree_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coruna-gophers/generics-poc/tree"
)

func TestInsert(t *testing.T) {
	tr := tree.New(compareInt)
	tr.Insert(1, "v1")
	tr.Insert(2, "v2")
	tr.Insert(3, "v3")
	tr.Insert(4, "v4")

	expected := map[int]string{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	}
	result := map[int]string{}
	tr.Walk(func(key, value interface{}) {
		result[key.(int)] = value.(string)
	})

	assert.Equal(t, expected, result)
}

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
