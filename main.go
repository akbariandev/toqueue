package main

import (
	"errors"
)

// Message struct contains Value of published message as any
type Message struct {
	Value any
}

// Channel struct contains a channel of Message struct
type Channel struct {
	ch chan Message
}

// Queue struct contains map of channels and sync.once for one time subscribe
type Queue struct {
	topics map[string]*Channel
}

// New will create and return an object of Queue
func New() *Queue {
	return &Queue{
		topics: map[string]*Channel{},
	}
}

// Subscribe on specific topic of queue. This method handle an action on specific message with handler func
func (q *Queue) Subscribe(topic string, process func(m *Message) bool, attempts int) error {

	if _, exist := q.topics[topic]; exist {
		return errors.New("subscribe exist on topic:" + topic)
	}

	q.topics[topic] = &Channel{
		ch: make(chan Message),
	}

	go func() {
		for {
			ch := <-q.topics[topic].ch
			if ok := process(&ch); !ok {
				if attempts == -1 || attempts > 0 {
					go func() {
						i := 0
						for i < attempts {
							if ok := process(&ch); ok {
								break
							}

							i++
							if attempts == -1 {
								i = 0
							}
						}
					}()
				}
			}
		}
	}()

	return nil
}

// Publish a message into specific topic. If topic not exist an error will return
func (q *Queue) Publish(topic string, msg Message) error {
	if _, ok := q.topics[topic]; !ok {
		return errors.New("topic not found")
	}

	q.topics[topic].ch <- msg

	return nil
}
