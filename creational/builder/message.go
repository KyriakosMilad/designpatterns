package main

import "time"

type msgType uint8

const (
	TextMsgType  msgType = 1
	VoiceMsgType msgType = 2
	ImageMsgType msgType = 3
	VideoMsgType msgType = 4
)

type Message struct {
	body      string
	typ       msgType
	createdAt time.Time
}
