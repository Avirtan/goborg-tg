package handlers

import (
	"context"

	message_dto "github.com/Avirtan/goborg-tg/dto/message"
	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	update_dto "github.com/Avirtan/goborg-tg/dto/update"
	"github.com/Avirtan/goborg-tg/method"
)

type TestCommandHandler struct {
}

func (c *TestCommandHandler) Action(ctx context.Context, update *update_dto.Update) {
	if update.Message != nil {
		method.SendMessage(
			ctx,
			method_dto.SendMessage[int64]{
				ChatID:      update.Message.From.Id,
				Text:        "test",
				ReplyMarkup: message_dto.Keyboard{
					// InlineKeyboard: dto.InlineKeyboard{
					// 	InlineKeyboard: [][]dto.InlineKeyboardButton{
					// 		{
					// 			{
					// 				Text:         "test",
					// 				CallbackData: "kb1",
					// 			},
					// 			{
					// 				Text:         "test1",
					// 				CallbackData: "kb2",
					// 			},
					// 		},
					// 		{
					// 			{
					// 				Text:         "test3",
					// 				CallbackData: "kb3",
					// 			},
					// 		},
					// 	}},
					// ForceReply: message_dto.ForceReply{
					// 	ForceReply:            true,
					// 	InputFieldPlaceholder: "test",
					// },
				},
			},
		)
	}
}
