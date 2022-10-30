package main

import "fmt"

func main() {
	roundedPizza := Pizza{}
	roundedPizza.Init(func() {
		fmt.Println("defining pizza shape to be rounded")
	})
	roundedPizza.MakePizza()

	square := Pizza{}
	square.Init(func() {
		fmt.Println("defining pizza shape to be square")
	})
	square.MakePizza()
}
