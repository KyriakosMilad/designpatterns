package main

import "fmt"

type Request struct {
	Amount uint
}

type Handler interface {
	SetSuccessor(h Handler)
	Handle(r *Request)
}

type ProjectManager struct {
	successor Handler
}

func (pm *ProjectManager) SetSuccessor(h Handler) {
	pm.successor = h
}

func (pm *ProjectManager) Handle(r *Request) {
	if r.Amount <= 100 {
		fmt.Printf("amount: %v handled by project manager\n", r.Amount)
	} else {
		fmt.Printf("amount: %v can't be handled by project manager, passed to successor\n", r.Amount)
		pm.successor.Handle(r)
	}
}

type CTO struct {
	successor Handler
}

func (ceo *CTO) SetSuccessor(h Handler) {
	ceo.successor = h
}

func (ceo *CTO) Handle(r *Request) {
	if r.Amount <= 1000 {
		fmt.Printf("amount: %v handled by cto\n", r.Amount)
	} else {
		fmt.Printf("amount: %v can't be handled by cto, passed to successor\n", r.Amount)
		ceo.successor.Handle(r)
	}
}

type CEO struct {
	successor Handler
}

func (ceo *CEO) SetSuccessor(h Handler) {
	ceo.successor = h
}

func (ceo *CEO) Handle(r *Request) {
	fmt.Printf("amount: %v handled by ceo\n", r.Amount)
}
