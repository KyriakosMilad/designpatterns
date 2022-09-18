package main

import (
	"fmt"
	"time"
)

type MessageBuilder struct {
	Body      string
	Typ       msgType
	CreatedAt time.Time
}

func newMessageBuilder() *MessageBuilder {
	return &MessageBuilder{}
}

func (mb *MessageBuilder) SetBody(body string) {
	mb.Body = body
}

func (mb *MessageBuilder) SetTyp(typ msgType) {
	mb.Typ = typ
}

func (mb *MessageBuilder) SetCreatedAt(createdAt time.Time) {
	mb.CreatedAt = createdAt
}

func (mb *MessageBuilder) Build() (*Message, error) {
	if mb.Typ != TextMsgType && mb.Typ != VoiceMsgType && mb.Typ != ImageMsgType && mb.Typ != VideoMsgType {
		return nil, fmt.Errorf("invalid message type: %v", mb.Typ)
	}

	return &Message{
		body:      mb.Body,
		typ:       mb.Typ,
		createdAt: mb.CreatedAt,
	}, nil
}
