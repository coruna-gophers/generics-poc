package tree_test

import (
	"fmt"
	"math/rand"
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

var sizes = []int{1e2, 1e3, 1e4}

func BenchmarkInsert(b *testing.B) {
	for _, n := range sizes {
		b.Run(fmt.Sprintf("BenchmarkInsert_%d", n), func(b *testing.B) {
			s := generateRandomSlice(n)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				getTree(s)
			}
		})
	}
}

func BenchmarkFind(b *testing.B) {
	for _, n := range sizes {
		b.Run(fmt.Sprintf("BenchmarkFind_%d", n), func(b *testing.B) {
			s := generateRandomSlice(n)
			tr := getTree(s)
			key := randRange(0, n)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tr.Find(key)
			}
		})
	}
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

func generateRandomSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = randRange(0, n)
	}
	return s
}

func getTree(s []int) *tree.Tree {
	tr := tree.New(compareInt)
	for k, v := range s {
		tr.Insert(k, v)
	}
	return tr
}

func randRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}
