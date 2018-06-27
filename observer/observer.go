package observer

import "fmt"

type emptyStruct struct{}

type Event interface {
}

type Observer interface {
	Update(Event)
}

type Observable interface {
	Add(Observer)
	Remove(Observer)
	Notify(Event)
}

type EventObserver struct {
	ID int
}

func (o *EventObserver) Update(e Event) {
	fmt.Printf("Observer %d received: %v\n", o.ID, e)
}

type eventObservable struct {
	observers map[Observer]emptyStruct
}

func New() Observable {
	return &eventObservable{
		observers: make(map[Observer]emptyStruct),
	}
}

func (c *eventObservable) Add(o Observer) {
	c.observers[o] = emptyStruct{}
}

func (c *eventObservable) Remove(o Observer) {
	delete(c.observers, o)
}

func (c *eventObservable) Notify(e Event) {
	for o := range c.observers {
		o.Update(e)
	}
}
