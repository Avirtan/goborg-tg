package handler

import (
	"context"

	update_dto "github.com/Avirtan/TGoBot/dto/update"
)

type TypeHandler uint64

const (
	Message TypeHandler = 1 << iota
	EditedMessage
	ChannelPost
	EditedChannelPost
)

type IHandler interface {
	Action(ctx context.Context, update *update_dto.Update)
}
