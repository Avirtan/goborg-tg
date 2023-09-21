package handlers

import (
	"TGoBot/dto"
	"TGoBot/method"
	"fmt"
)

type TestCommandHandler struct {
}

func (c *TestCommandHandler) Action(update *dto.Update, msgHandler *method.MethodHandler) {
	if update.Message != nil {
		fmt.Println(update.Message.MessageId)
		msgHandler.SendMessage(
			dto.SendMessage{
				ChatID: update.Message.From.Id,
				Text:   "test",
				// ReplyMarkup: dto.ForceReply{
				// 	ForceReply: true,
				// },
			},
		)
	}
}
