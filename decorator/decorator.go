package decorator

type Component interface {
	Action() int
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Action() int {
	return 0
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

type AddDecorator struct {
	Component
	num int
}

func (d *AddDecorator) Action() int {
	return d.Component.Action() + d.num
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

type MulDecorator struct {
	Component
	num int
}

func (d *MulDecorator) Action() int {
	return d.Component.Action() * d.num
}
