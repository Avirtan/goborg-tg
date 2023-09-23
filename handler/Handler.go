package handler

import (
	"TGoBot/dto"
	"TGoBot/method"
	"context"
)

type TypeHandler uint64

const (
	Message TypeHandler = 1 << iota
	EditedMessage
	ChannelPost
	EditedChannelPost
)

type IHandler interface {
	Action(ctx context.Context, update *dto.Update, msgHandler *method.MethodHandler)
}
