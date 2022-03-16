package basic

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func print[T any](x ...T) {
	fmt.Println(x)
}

func isOrdered[T constraints.Ordered](x, y T) bool {
	return x < y
}
