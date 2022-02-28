package kv_test

import (
	"fmt"
	"github.com/coruna-gophers/generics-poc/kv"
	"strconv"
	"testing"
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
