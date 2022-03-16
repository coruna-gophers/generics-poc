package list_test

import (
	"fmt"
	"testing"

	"github.com/coruna-gophers/generics-poc/list"
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

	var findOdd list.FindFn[int]
	findOdd = func(index int, item int) bool {
		return item%2 == 1
	}

	index, err := l.FindWhere(findOdd)
	if err != nil {
		t.Fatalf("error finding first even element: %v", err)
	}
	fmt.Printf("first event element at %d\n", index)
}
