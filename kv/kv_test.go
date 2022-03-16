package kv_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/coruna-gophers/generics-poc/kv"
)

func TestKV(t *testing.T) {
	kv := kv.New[string, int]()
	kv.Put("foo", 1)
	kv.Put("bar", 2)

	v, err := kv.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println("Key foo, Value ", strconv.Itoa(v))
	v, err = kv.Get("bar")
	if err != nil {
		panic(err)
	}
	fmt.Println("Key bar, Value ", strconv.Itoa(v))
}
