package main

import (
	"fmt"
	"os"
	"time"

	pubisubo "github.com/pedrobertao/pubisubo/pubsub"
)

func main() {
	fmt.Println("Welcome to PubiSubo")
	pbStr := pubisubo.Start[string]()
	pbFloat := pubisubo.Start[float64]()

	msg1 := "First msg"
	msg2 := "Second msg"

	msg3 := 3.14159
	msg4 := 2.71828

	topic := "TEST_STRING"
	topic2 := "TEST_FLOAT"

	c1 := pbStr.Subscribe(topic)
	c2 := pbStr.Subscribe(topic)

	c3 := pbFloat.Subscribe(topic2)
	c4 := pbFloat.Subscribe(topic2)

	chEnd := make(chan bool)
	go func() {
		for msg := range c1 {
			fmt.Printf("[C1]Message receive from %s: %+v\n", topic, msg)

			// This should be adapted, but was write for demo purpose
			if msg == msg2 {
				chEnd <- true
				break
			}
		}
	}()

	go func() {
		for msg := range c2 {
			fmt.Printf("[C2]Message receive from %s: %+v\n", topic, msg)

			// This should be adapted, but was write for demo purpose
			if msg == msg2 {
				chEnd <- true
				break
			}
		}
	}()

	go func() {
		for msg := range c3 {
			fmt.Printf("[C3]Message receive from %s: %+v\n", topic, msg)

			// This should be adapted, but was write for demo purpose
			if msg == msg4 {
				chEnd <- true
				break
			}
		}
	}()

	go func() {
		for msg := range c4 {
			fmt.Printf("[C4]Message receive from %s: %+v\n", topic, msg)

			// This should be adapted, but was write for demo purpose
			if msg == msg4 {
				chEnd <- true
				break
			}
		}
	}()

	// Publishing messages
	pbStr.Publish(topic, msg1)
	pbStr.Publish(topic, msg2)

	time.Sleep(3 * time.Second)

	pbFloat.Publish(topic2, msg3)
	pbFloat.Publish(topic2, msg4)

	count := 0
	for msg := range chEnd {
		if msg {
			count++
		}

		if count == 4 {
			fmt.Println("Channels closed")
			close(c1)
			close(c2)
			close(c3)
			close(c4)

			fmt.Println("All done !")
			os.Exit(0)
		}
	}
}
