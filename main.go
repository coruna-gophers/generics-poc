package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"

	"github.com/coruna-gophers/generics-poc/channel"
	"github.com/coruna-gophers/generics-poc/kv"
)

func print[T any](x ...T) {
	fmt.Println(x)
}

func isOrdered[T constraints.Ordered](x, y T) bool {
	return x < y
}

func consume[T any](id string, interval time.Duration, fn channel.Pull[T]) {
	for {
		elem, err := fn()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("rcv %s: %v \n", id, elem)
		time.Sleep(interval)
	}
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

	fo := channel.NewFanOut[int]()

	pull1, unsubscribe1 := fo.Subscribe(3)
	defer unsubscribe1()
	go consume[int]("ch1", time.Second, pull1)

	pull2, unsubscribe2 := fo.Subscribe(3)
	defer unsubscribe2()
	go consume[int]("ch2", 2*time.Second, pull2)

	for i := 0; i < 10; i++ {
		fo.Publish(i)
	}
	time.Sleep(14 * time.Second)
}
