package main

import (
	"fmt"
	"reflect"
)

func main() {
	o := Object{
		x: 55,
		Y: 99,
	}

	oClone := o.Clone()

	fmt.Println(reflect.DeepEqual(o, oClone)) // should print true
}
