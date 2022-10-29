package main

import (
	"fmt"
	"reflect"
)

func main() {
	gf := &GenFactory{}

	btn := gf.createBtn(MacOSType)
	fmt.Println(reflect.ValueOf(btn).Type()) // should print MacBtn

	btn = gf.createBtn(WinOSType)
	fmt.Println(reflect.ValueOf(btn).Type()) // should print WinBtn
}
