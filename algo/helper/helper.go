package helper

import (
	"math/rand"

	"golang.org/x/exp/slices"
)

func Higher(s []int) int {
	scopy := make([]int, len(s))
	copy(scopy, s)
	slices.Sort(scopy)
	return scopy[len(s)-1]
}

func GenerateRandomSliceSet(n int) []int {
	s := make([]int, n)
	set := map[int]struct{}{}
	for i := 0; i < n; i++ {
		rn := RandRange(0, n)
		_, ok := set[rn]
		if ok {
			i--
			continue
		}
		set[rn] = struct{}{}
		s[i] = rn
	}
	return s
}

func RandRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}
