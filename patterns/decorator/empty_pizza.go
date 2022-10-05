package decorator

type EmptyPizza struct {
}

func (p EmptyPizza) GetPrice() int {
	return 15
}
