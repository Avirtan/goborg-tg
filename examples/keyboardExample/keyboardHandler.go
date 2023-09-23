package main

import (
	"TGoBot/dto"
	"TGoBot/method"
	"context"
)

type TestCommandHandler struct {
}

func (c *TestCommandHandler) Action(ctx context.Context, update *dto.Update, msgHandler *method.MethodHandler) {
	if update.Message != nil {
		msgHandler.SendMessage(
			dto.SendMessage{
				ChatID: update.Message.From.Id,
				Text:   "test",
				ReplyMarkup: dto.Keyboard{
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
					ForceReply: dto.ForceReply{
						ForceReply:            true,
						InputFieldPlaceholder: "test",
					},
				},
			},
		)
	}
}
