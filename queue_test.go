package main

import "testing"

var tests = []struct {
	Topic   string
	Message any
}{
	{"Topic 1", "Message 1"},
	{"Topic 2", 123},
	{"Topic 3", struct{}{}},
	{"Topic 4", true},
}

func TestQueue_Publish(t *testing.T) {

	q := New()
	for _, test := range tests {
		if err := q.Publish(test.Topic, Message{test.Message}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestQueue_Subscribe(t *testing.T) {
	q := New()
	for _, test := range tests {
		if err := q.Subscribe(test.Topic, processMessage, -1); err != nil {
			t.Fatal(err)
		}
	}
}

func processMessage(m *Message) bool {
	return true
}
