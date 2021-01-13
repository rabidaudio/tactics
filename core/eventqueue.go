package core

import (
	uuid "github.com/hashicorp/go-uuid"
	"github.com/rabidaudio/tactics/core/units"
)

var Events = NewQueue()

type Event interface {
	// Action() string
}

type Handler func(event Event, tick units.Tick)

type EventQueue struct {
	queue    []scheduledEvent
	handlers map[string]Handler
	tick     units.Tick
}

func NewQueue() EventQueue {
	return EventQueue{
		queue:    make([]scheduledEvent, 0, 25),
		handlers: make(map[string]Handler),
		tick:     0,
	}
}

type scheduledEvent struct {
	event Event
	at    units.Tick
}

func (eq *EventQueue) Tick() {
	eq.tick++
	handled := 0
	for _, se := range eq.queue {
		if se.at > eq.tick {
			break
		}
		for _, h := range eq.handlers {
			h(se.event, eq.tick)
		}
		handled++
	}
	eq.queue = eq.queue[handled:]
}

func (eq *EventQueue) AddListener(handler Handler) string {
	key, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}
	eq.handlers[key] = handler
	return key
}

func (eq *EventQueue) RemoveListener(key string) {
	delete(eq.handlers, key)
}

func (eq *EventQueue) Schedule(event Event, in units.Tick) {
	if in < 0 {
		panic("Cannot schedule events in the past")
	}
	se := scheduledEvent{event: event, at: eq.tick + in}
	// shortcut for the most common cases
	if len(eq.queue) == 0 || se.at >= eq.queue[len(eq.queue)-1].at {
		eq.queue = append(eq.queue, se)
		return
	}
	if eq.queue[0].at > se.at {
		eq.queue = append([]scheduledEvent{se}, eq.queue...)
		return
	}
	for i := len(eq.queue) - 2; i >= 0; i-- {
		if eq.queue[i].at <= se.at {
			eq.queue = append(eq.queue[:i+2], eq.queue[i+1:]...)
			eq.queue[i+1] = se
			return
		}
	}
}
