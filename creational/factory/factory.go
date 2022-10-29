package main

import (
	"fmt"
)

const (
	TunaPizzaType       = "TunaPizza"
	CheesePizzaType     = "CheesePizza"
	VegetablesPizzaType = "VegetablesPizza"
)

type Pizza interface{}

type TunaPizza struct{}
type CheesePizza struct{}
type VegetablesPizza struct{}

type Chef struct{}

func (c *Chef) MakePizza(pizzaType string) (error, Pizza) {
	if pizzaType == TunaPizzaType {
		return nil, &TunaPizza{}
	} else if pizzaType == CheesePizzaType {
		return nil, &CheesePizza{}
	} else if pizzaType == VegetablesPizzaType {
		return nil, &VegetablesPizza{}
	}

	return fmt.Errorf("pizza type `%v` unknown", pizzaType), nil
}
