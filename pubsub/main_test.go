package pubsub

import (
	"testing"
)

func TestPubSubString(t *testing.T) {
	pbStr := Start[string]()

	ch := pbStr.Subscribe("test")
	chEnd := make(chan bool)

	go func() {
		for msg := range ch {
			if msg == "test" {
				chEnd <- true
				break
			}
		}
	}()

	pbStr.Publish("test", "test")

	go func() {
		for msg := range chEnd {
			if msg {
				break
			}
		}
	}()
}

func TestPubSubFloat(t *testing.T) {
	pbStr := Start[float64]()
	numTest := 3.141

	ch := pbStr.Subscribe("test")
	chEnd := make(chan bool)

	go func() {
		for msg := range ch {
			if msg == numTest {
				chEnd <- true
				break
			}
		}
	}()

	pbStr.Publish("test", numTest)

	go func() {
		for msg := range chEnd {
			if msg {
				break
			}
		}
	}()
}
