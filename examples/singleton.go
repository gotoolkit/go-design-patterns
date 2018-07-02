package main

import (
	"github.com/gotoolkit/go-design-patterns/singleton"
)

func main() {
	singleton.Instance().Count()
	singleton.Instance().Count()
	singleton.Instance().Count()

	s := singleton.Instance()
	s.Count()

	s1 := singleton.Instance()
	s1.Count()

	s2 := singleton.Instance()
	s2.Count()
}
