package main

import "reflect"

type Object struct {
	singleInstance *Object
}

func (o *Object) GetInstance() *Object {
	if reflect.ValueOf(o.singleInstance).IsNil() {
		o.singleInstance = &Object{}
	}

	return o.singleInstance
}
