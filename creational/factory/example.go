package main

import (
	"fmt"
	"reflect"
)

func main() {
	c := &Chef{}

	err, p := c.MakePizza(TunaPizzaType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reflect.ValueOf(p).Type())
	}

	err, p = c.MakePizza(CheesePizzaType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reflect.ValueOf(p).Type())
	}

	err, p = c.MakePizza(VegetablesPizzaType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reflect.ValueOf(p).Type())
	}

	err, p = c.MakePizza("Burger")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reflect.ValueOf(p).Type())
	}
}
