package main

import "learn/patterns/singleton"

func main() {
	// NOTE: decorator
	//pizza := decorator.EmptyPizza{}
	//
	//fmt.Println("Raw pizza price", pizza.GetPrice())
	//
	//cheesePizza := decorator.CheesePizza{
	//	Pizza: pizza,
	//}
	//
	//fmt.Println("Cheese pizza price", cheesePizza.GetPrice())
	//
	//beefAndCheesePizza := decorator.BeefPizza{
	//	Pizza: cheesePizza,
	//}
	//
	//fmt.Println("Cheese pizza price", beefAndCheesePizza.GetPrice())

	// NOTE: proxy
	//proxyServer := proxy.InitNginx()
	//
	//code, msg := proxyServer.HandleRequest("/user", "get")
	//
	//fmt.Println("code:", code, ", msg:", msg)
	//
	//code, msg = proxyServer.HandleRequest("/user", "get")
	//fmt.Println("code:", code, ", msg:", msg)
	//
	//code, msg = proxyServer.HandleRequest("/user", "get")
	//fmt.Println("code:", code, ", msg:", msg)

	//NOTE: strategy

	//addStrategy := strategy.AddStrategy{}
	//
	//context := strategy.InitContext(addStrategy)
	//
	//minusStrategy := strategy.MinusStrategy{}
	//
	//context.SetStrategy(minusStrategy)
	//
	//fmt.Println(context.Execute(1, 2))

	//NOTE: singleton
	singleton.Execute()
}
