package event

import "sync"

type EventChannel struct {
	C      chan Event
	Locker sync.RWMutex
}

type Event struct {
	Type   string
	Object interface{}
}

func NewEventChan() *EventChannel {
	return &EventChannel{
		C: make(chan Event, 128),
	}
}

func NewEvent() {
}
