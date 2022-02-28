package channel_test

import (
	"fmt"
	"github.com/coruna-gophers/generics-poc/channel"
	"io"
	"log"
	"testing"
	"time"
)

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

func TestNewFanOut(t *testing.T) {
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
