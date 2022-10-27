package main

import "fmt"

func main() {
	c := &Caretaker{}

	g := CreateGame("kyriakos", 0)

	c.Save(g, 9)

	g.SetScore(99)

	fmt.Printf("before undo --> %v\n", g.GetScore()) // should print 99

	c.Revert(g, 9)

	fmt.Printf("after undo --> %v\n", g.GetScore()) // should print 0
}
