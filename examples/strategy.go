package main

import (
	"github.com/gotoolkit/go-design-patterns/strategy"
)

func main() {
	ctx := strategy.NewContext("Paul", &strategy.DefaultStrategy{})
	ctx.Action()

	ctx = strategy.NewContext("Bob", &strategy.VariedStrategy{})
	ctx.Action()
}
