package channel

import (
	"fmt"
	"io"

	"github.com/google/uuid"
)

type GenericChan[T any] chan T

type FanOut[T any] struct {
	channels map[string]GenericChan[T]
}

func NewFanOut[T any]() *FanOut[T] {
	return &FanOut[T]{channels: map[string]GenericChan[T]{}}
}

type Pull[T any] func() (T, error)
type Unsubscribe func() error

func (fo *FanOut[T]) Subscribe(size int) (pull Pull[T], unsubs Unsubscribe) {
	uid := uuid.New().String()
	ch := make(GenericChan[T], size)
	fo.channels[uid] = ch
	unsubs = func() error {
		if _, ok := fo.channels[uid]; !ok {
			return fmt.Errorf("channel uid: %s not found", uid)
		}
		delete(fo.channels, uid)
		return nil
	}
	pull = func() (value T, err error) {
		value, ok := <-ch
		if !ok {
			err = io.EOF
		}
		return
	}
	return
}

func (fo *FanOut[T]) Publish(elem T) {
	for i := range fo.channels {
		fo.channels[i] <- elem
	}
}
