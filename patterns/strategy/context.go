package strategy

type Context struct {
	strategy IStrategy
}

func InitContext(s IStrategy) *Context {
	return &Context{strategy: s}
}

func (c *Context) SetStrategy(s IStrategy) {
	c.strategy = s
}

func (c *Context) Execute(a int, b int) int {
	return c.strategy.Execute(a, b)
}
