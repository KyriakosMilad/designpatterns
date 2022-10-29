package main

import "fmt"

func main() {
	mathBook := NewBook("math", nil)
	historyBook := NewBook("history", mathBook)
	scienceBook := NewBook("science", historyBook)
	algorithmsBook := NewBook("algorithms", scienceBook)

	w := NewWorker(algorithmsBook)

	fmt.Println(w.getNextBook())
	fmt.Println(w.getNextBook())
	fmt.Println(w.getNextBook())
	fmt.Println(w.getNextBook())
}
