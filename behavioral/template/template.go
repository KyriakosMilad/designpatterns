package main

import "fmt"

type Pizza struct {
	definePizzaShape func()
}

func (p *Pizza) Init(definePizzaShape func()) {
	p.definePizzaShape = definePizzaShape
}

func (p *Pizza) bakePizza() {
	fmt.Println("backing the pizza")
}

func (p *Pizza) putAddons() {
	fmt.Println("putting addons on the pizza")
}

func (p *Pizza) heatPizza() {
	fmt.Println("heating the pizza")
}

func (p *Pizza) MakePizza() {
	p.bakePizza()
	p.definePizzaShape()
	p.putAddons()
	p.heatPizza()
}
