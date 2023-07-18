package pubsub

import (
	"fmt"
	"sync"
)

type GenericValues interface {
	string | float64
}

type PubSub[T GenericValues] struct {
	subs map[string][]chan T
	mu   sync.Mutex
}

func (p *PubSub[T]) Subscribe(topic string) chan T {
	fmt.Printf("Subscribing for topic %s\n", topic)

	p.mu.Lock()
	defer p.mu.Unlock()

	ch := make(chan T)
	p.subs[topic] = append(p.subs[topic], ch)

	return ch
}

func (p *PubSub[T]) Publish(topic string, msg T) {
	fmt.Printf("Sending Message to topic: %s\n", topic)

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.subs[topic] != nil {
		for _, chs := range p.subs[topic] {
			chs <- msg
		}
	} else {
		fmt.Println("Topic not available")
	}
}

func Start[T GenericValues]() *PubSub[T] {
	fmt.Println("*=Starting PubSub=*")
	pg := &PubSub[T]{
		subs: make(map[string][]chan T),
		mu:   sync.Mutex{},
	}
	return pg
}
