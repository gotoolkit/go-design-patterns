package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
	count int
}

var instance *singleton
var once sync.Once

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) Count() {
	s.count++
	fmt.Println(s.count)
}
