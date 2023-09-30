package main

import (
	message_dto "TGoBot/dto/message"
	method_dto "TGoBot/dto/method"
	update_dto "TGoBot/dto/update"
	"TGoBot/method"
	"context"
)

type TestCommandHandler struct {
}

func (c *TestCommandHandler) Action(ctx context.Context, update *update_dto.Update, msgHandler *method.MethodHandler) {
	if update.Message != nil {
		msgHandler.SendMessage(
			method_dto.SendMessage{
				ChatID: update.Message.From.Id,
				Text:   "test",
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
					ForceReply: message_dto.ForceReply{
						ForceReply:            true,
						InputFieldPlaceholder: "test",
					},
				},
			},
		)
	}
}
