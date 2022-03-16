package kv

import "errors"

type Store[K comparable, V any] struct {
	store map[K]V
}

func (kv *Store[K, V]) Get(key K) (result V, err error) {
	if v, ok := kv.store[key]; ok {
		result = v
		return
	}
	err = errors.New("not found")
	return
}

func (kv *Store[K, V]) Put(key K, value V) error {
	kv.store[key] = value
	return nil
}

func New[K comparable, V any]() *Store[K, V] {
	return &Store[K, V]{
		store: make(map[K]V),
	}
}
