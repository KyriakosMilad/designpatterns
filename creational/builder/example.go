package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bldr := NewMessageBuilder()

	bldr.SetBody("hello")
	bldr.SetTyp(TextMsgType)
	//bldr.SetTyp(66) // uncomment to test handling error
	bldr.SetCreatedAt(time.Now())

	msg, err := bldr.Build()
	if err != nil {
		log.Fatalf("error building message: %v", err)
	} else {
		fmt.Printf("%v", msg)
	}
}
