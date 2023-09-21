package handler

import (
	"TGoBot/dto"
	"TGoBot/method"
)

type TypeHandler uint64

const (
	Message TypeHandler = 1 << iota
	EditedMessage
	ChannelPost
	EditedChannelPost
)

type IHandler interface {
	Action(update *dto.Update, msgHandler *method.MethodHandler)
}
