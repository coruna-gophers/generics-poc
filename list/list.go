package list

import (
	"errors"
	"fmt"
)

type List[T any] struct {
	arr []T
}

func (l *List[T]) Add(item ...T) {
	l.arr = append(l.arr, item...)
}

type FindFn[T any] func(index int, item T) bool

func (l *List[T]) FindWhere(fn FindFn[T]) (int, error) {
	for i, v := range l.arr {
		if fn(i, v) {
			return i, nil
		}
	}
	return -1, nil
}

func (l *List[T]) Remove(index int) error {
	if err := l.checkIndex(index); err != nil {
		return err
	}
	l.arr = append(l.arr[:index], l.arr[index+1:]...)
	return nil
}

func (l *List[T]) Update(index int, value T) error {
	if err := l.checkIndex(index); err != nil {
		return err
	}
	l.arr[index] = value
	return nil
}

func (l *List[T]) checkIndex(index int) error {
	if index < 0 || index > (len(l.arr)-1) {
		return errors.New("index out of bounds")
	}
	return nil
}

func (l *List[T]) String() string {
	return fmt.Sprintf("%v len=%d cap=%d", l.arr, len(l.arr), cap(l.arr))
}

func NewList[T any]() *List[T] {
	return &List[T]{
		arr: []T{},
	}
}
