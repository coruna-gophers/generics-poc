package gtree_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"

	"github.com/coruna-gophers/generics-poc/algo/gtree"
	"github.com/coruna-gophers/generics-poc/algo/helper"
)

func TestTree_Insert(t *testing.T) {
	tr := gtree.New[int, string](compare[int])
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
	tr.Walk(func(key int, value string) {
		result[key] = value
	})
	assert.Equal(t, expected, result)
}

func compare[K constraints.Ordered](keyA, keyB K) int {
	if keyA == keyB {
		return 0
	}
	if keyA < keyB {
		return -1
	}
	if keyA == keyB {
		return 0
	}
	return 1
}

func TestTree_Find(t *testing.T) {
	tr := gtree.New[int, string](compare[int])
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
	tr := gtree.New[int, string](compare[int])
	tr.Insert(1, "v1")
	tr.Insert(15, "v15")
	tr.Insert(121, "v121")
	tr.Insert(5657, "v5657")
	tr.Insert(12, "v12")
	tr.Insert(0, "v0")

	value, found := tr.Find(666)
	require.False(t, found)
	require.Empty(t, value)
}

var sizes = []int{1e2, 1e3, 1e4, 1e6}

func BenchmarkInsert(b *testing.B) {
	for _, n := range sizes {
		b.Run(fmt.Sprintf("BenchmarkInsert_%d", n), func(b *testing.B) {
			s := helper.GenerateRandomSliceSet(n)
			key := helper.Higher(s) + 1
			tr := getTree(s)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tr.Insert(key, key)
			}
		})
	}
}

var findOut interface{}
var findFoundOut bool

func BenchmarkFind(b *testing.B) {
	for _, n := range sizes {
		b.Run(fmt.Sprintf("BenchmarkFind_%d", n), func(b *testing.B) {
			s := helper.GenerateRandomSliceSet(n)
			key := helper.Higher(s)
			tr := getTree(s)

			var o1 interface{}
			var o2 bool

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				o1, o2 = tr.Find(key)
			}

			findOut = o1
			findFoundOut = o2
		})
	}
}

func getTree(s []int) *gtree.Tree[int, int] {
	tr := gtree.New[int, int](compare[int])
	for _, v := range s {
		tr.Insert(v, v)
	}
	return tr
}
