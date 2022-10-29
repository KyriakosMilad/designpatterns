package main

import (
	"fmt"
	"reflect"
)

func main() {
	o := &Object{}

	i1 := o.GetInstance()
	i2 := o.GetInstance()

	fmt.Println(reflect.DeepEqual(i1, i2)) // should print true
}
