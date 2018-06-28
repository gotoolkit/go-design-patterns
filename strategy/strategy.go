package strategy

import "fmt"

type Context struct {
	Name     string
	strategy Strategy
}

func NewContext(name string, strategy Strategy) *Context {
	return &Context{
		Name:     name,
		strategy: strategy,
	}
}

func (ctx *Context) Action() {
	ctx.strategy.Action(ctx)
}

type Strategy interface {
	Action(*Context)
}

type DefaultStrategy struct {
}

func (d *DefaultStrategy) Action(ctx *Context) {
	fmt.Printf("Action %s by Default\n", ctx.Name)
}

type VariedStrategy struct {
}

func (d *VariedStrategy) Action(ctx *Context) {
	fmt.Printf("Action %s by Variation\n", ctx.Name)
}
