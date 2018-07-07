package main

import (
	"github.com/gotoolkit/go-design-patterns/strategy"
)

func main() {
	
	ctx := strategy.New("Paul", nil)
	ctx.Run()
	
	strategy.Start(&strategy.VariedStrategy{})
}
