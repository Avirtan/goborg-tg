package method_dto

import message_dto "TGoBot/dto/message"

type SendMessage struct {
	ChatID      int64                `json:"chat_id"`
	Text        string               `json:"text"`
	ReplyMarkup message_dto.Keyboard `json:"reply_markup,omitempty"`
}
