package main

type Object struct {
	x int
	Y int
}

func (o *Object) Clone() Object {
	return Object{x: o.x, Y: o.Y}
}
