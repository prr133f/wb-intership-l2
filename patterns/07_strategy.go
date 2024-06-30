package patterns

import "fmt"

type IStrategy interface {
	execute(a, b int) int
}

type Context struct {
	strategy IStrategy
}

func NewContext(strategy IStrategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (c *Context) SetStrategy(strategy IStrategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.execute(a, b)
}

type Division struct {
	a int
	b int
}

func NewDivision(a, b int) *Division {
	return &Division{
		a: a,
		b: b,
	}
}

func (d *Division) execute(a, b int) int {
	return a / b
}

type Multiplication struct {
	a int
	b int
}

func NewMultiplication(a, b int) *Multiplication {
	return &Multiplication{
		a: a,
		b: b,
	}
}

func (m *Multiplication) execute(a, b int) int {
	return a * b
}

type Addition struct {
	a int
	b int
}

func NewAddition(a, b int) *Addition {
	return &Addition{
		a: a,
		b: b,
	}
}

func (add *Addition) execute(a, b int) int {
	return a + b
}

type Subtraction struct {
	a int
	b int
}

func NewSubtraction(a, b int) *Subtraction {
	return &Subtraction{
		a: a,
		b: b,
	}
}

func (sub *Subtraction) execute(a, b int) int {
	return a - b
}

func strategy() {
	add := NewAddition(1, 2)
	sub := NewSubtraction(1, 2)
	mul := NewMultiplication(1, 2)
	div := NewDivision(1, 2)

	c := &Context{}

	c.SetStrategy(add)
	fmt.Println(c.ExecuteStrategy(1, 2))

	c.SetStrategy(sub)
	fmt.Println(c.ExecuteStrategy(1, 2))

	c.SetStrategy(mul)
	fmt.Println(c.ExecuteStrategy(1, 2))

	c.SetStrategy(div)
	fmt.Println(c.ExecuteStrategy(1, 2))
}
