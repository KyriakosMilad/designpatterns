package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	name     string
	nextBook *Book
}

func NewBook(name string, nextBook *Book) *Book {
	return &Book{name: name, nextBook: nextBook}
}

type Worker struct {
	currentBook *Book
}

func NewWorker(book *Book) *Worker {
	return &Worker{currentBook: book}
}

func (w *Worker) getNextBook() string {
	if reflect.ValueOf(w.currentBook.nextBook).IsNil() {
		fmt.Println("last book reached")
		return ""
	}
	nextBook := w.currentBook.nextBook
	w.currentBook = nextBook

	return nextBook.name
}
