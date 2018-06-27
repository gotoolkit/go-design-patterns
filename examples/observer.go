package main

import (
	"time"

	"github.com/gotoolkit/go-design-patterns/observer"
)

func main() {
	l := observer.New()
	o3 := &observer.EventObserver{3}
	l.Add(&observer.EventObserver{1})
	l.Add(&observer.EventObserver{2})
	l.Add(o3)
	l.Add(&observer.EventObserver{4})
	l.Remove(o3)
	l.Notify(time.Now().UnixNano())
}
