package decorator

type BeefPizza struct {
	Pizza IPizza
}

func (bp BeefPizza) GetPrice() int {
	price := bp.Pizza.GetPrice()
	return price + 20
}
