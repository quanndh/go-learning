package strategy

type AddStrategy struct {
}

func (s AddStrategy) Execute(a int, b int) int {
	return a + b
}

type MinusStrategy struct {
}

func (s MinusStrategy) Execute(a int, b int) int {
	return a - b
}
