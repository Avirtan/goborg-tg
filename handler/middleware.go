package handler

import (
	"context"

	update_dto "github.com/Avirtan/TGoBot/dto/update"
)

type Middleware interface {
	Check(ctx context.Context, update *update_dto.Update) bool
}
