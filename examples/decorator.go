package main

import (
	"fmt"

	"github.com/gotoolkit/go-design-patterns/decorator"
)

func main() {

	var c decorator.Component = &decorator.ConcreteComponent{}
	c = decorator.WrapAddDecorator(c, 33)
	c = decorator.WrapMulDecorator(c, 3)
	res := c.Action()
	fmt.Printf("res: %d\n", res)
}
