package list_test

import (
	"fmt"
	"github.com/coruna-gophers/generics-poc/list"
	"testing"
)

func TestList(t *testing.T) {
	l := list.NewList[int]()
	fmt.Println(l)

	l.Add(1, 2, 3)
	fmt.Println(l)

	l.Remove(0)
	fmt.Println(l)

	l.Add(4)
	fmt.Println(l)

	l.Add(5, 6)
	fmt.Println(l)

	l.Update(3, 10)
	fmt.Println(l)

	var findFn list.FindFn[int]
	findFn = func(index int, item int) bool {
		return item%2 == 0
	}

	index, err := l.FindWhere(findFn)
	if err != nil {
		t.Fatalf("error finding first even element: %v", err)
	}
	fmt.Printf("first event element at %d\n", index)
}
