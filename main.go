package main

import (
	"fmt"
	"github.com/coruna-gophers/generics-poc/kv"
	"golang.org/x/exp/constraints"
	"strconv"
)

func print[T any](x ...T) {
	fmt.Println(x)
}

func isOrdered[T constraints.Ordered](x, y T) bool {
	return x < y
}

func main() {
	print([]int{1, 2, 3})
	print(isOrdered(1, 2))

	kv := kv.New[string, int]()
	kv.Put("foo", 1)
	kv.Put("bar", 2)

	v, err := kv.Get("foo")
	if err != nil {
		panic(err)
	}
	print("Key foo, Value ", strconv.Itoa(v))
	v, err = kv.Get("bar")
	if err != nil {
		panic(err)
	}
	print("Key bar, Value ", strconv.Itoa(v))
}
