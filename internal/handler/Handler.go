package handler

import (
	"TGoBot/internal/dto"
	"TGoBot/internal/method"
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
