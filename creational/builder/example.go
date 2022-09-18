package main

import (
	"fmt"
	"time"
)

func main() {
	bldr := newMessageBuilder()

	bldr.SetBody("hello")
	bldr.SetTyp(TextMsgType)
	//bldr.SetTyp(66) // uncomment to test handling error
	bldr.SetCreatedAt(time.Now())

	msg, err := bldr.Build()
	if err != nil {
		_ = fmt.Errorf("error building message: %v", err)
	} else {
		fmt.Printf("%v", msg)
	}
}
