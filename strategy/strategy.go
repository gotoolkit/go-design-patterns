package strategy

import "fmt"

type Context struct {
	Strategy
	Name     string
}

func New(name string, s Strategy) *Context {
	if s == nil {
		s = &DefaultStrategy{}
	}
	return &Context{
		Strategy: s,
		Name:     name,
	}
}

func Start(s Strategy) {
	s.Action()
}

func (ctx *Context) Run() {
	ctx.Action()
}

type Strategy interface {
	Action()
}

type DefaultStrategy struct {
}

func (d *DefaultStrategy) Action() {
	fmt.Println("Action by Default")
}

type VariedStrategy struct {
}

func (d *VariedStrategy) Action() {
	fmt.Println("Action by Variation")
}
