package handler

import (
	update_dto "TGoBot/dto/update"
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
	Action(ctx context.Context, update *update_dto.Update, msgHandler *method.MethodHandler)
}
