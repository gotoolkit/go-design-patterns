package strategy

func ExampleDefaultStrategy() {
	ctx := NewContext("Paul", &DefaultStrategy{})
	ctx.Action()
}

func ExampleVariedStrategy() {
	ctx := NewContext("Bob", &VariedStrategy{})
	ctx.Action()
}
