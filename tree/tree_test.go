package tree_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/coruna-gophers/generics-poc/tree"
)

func TestTree_Insert(t *testing.T) {
	tr := tree.New(compareInt)
	tr.Insert(4, "v4")
	tr.Insert(1, "v1")
	tr.Insert(3, "v3")
	tr.Insert(2, "v2")

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

func TestTree_Find(t *testing.T) {
	tr := tree.New(compareInt)
	tr.Insert(1, "v1")
	tr.Insert(15, "v15")
	tr.Insert(121, "v121")
	tr.Insert(5657, "v5657")
	tr.Insert(12, "v12")
	tr.Insert(0, "v0")

	value, found := tr.Find(12)
	require.True(t, found)
	require.Equal(t, "v12", value)
}

func TestTree_FindNotFound(t *testing.T) {
	tr := tree.New(compareInt)
	tr.Insert(1, "v1")
	tr.Insert(15, "v15")
	tr.Insert(121, "v121")
	tr.Insert(5657, "v5657")
	tr.Insert(12, "v12")
	tr.Insert(0, "v0")

	value, found := tr.Find(666)
	require.False(t, found)
	require.Nil(t, value)
}

func compareInt(keyA, keyB interface{}) int {
	if keyA == keyB {
		return 0
	}
	aInt, ok := keyA.(int)
	if !ok {
		panic(fmt.Errorf("'%s' is not an integer", keyA))
	}
	bInt, ok := keyB.(int)
	if !ok {
		panic(fmt.Errorf("'%s' is not an integer", keyB))
	}
	if aInt < bInt {
		return -1
	}
	if aInt == bInt {
		return 0
	}
	return 1
}
