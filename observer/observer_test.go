package observer

import "time"

func ExampleObserver() {
	l := New()
	l.Add(&EventObserver{1})
	l.Add(&EventObserver{2})
	l.Add(&EventObserver{3})
	l.Add(&EventObserver{4})
	l.Remove(&EventObserver{3})
	l.Notify(time.Now().UnixNano())
}
