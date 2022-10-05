package decorator

type CheesePizza struct {
	Pizza IPizza
}

func (cp CheesePizza) GetPrice() int {
	price := cp.Pizza.GetPrice()
	return price + 10
}
