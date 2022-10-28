package main

func main() {
	pm := &ProjectManager{}
	cto := &CTO{}
	ceo := &CEO{}

	cto.SetSuccessor(ceo)
	pm.SetSuccessor(cto)

	pm.Handle(&Request{Amount: 50})
	pm.Handle(&Request{Amount: 900})
	pm.Handle(&Request{Amount: 1200})
}
