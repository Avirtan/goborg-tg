package method_dto

import message_dto "TGoBot/dto/message"

type SendMessage struct {
	ChatID                   int64                       `json:"chat_id"`
	Text                     string                      `json:"text"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	Entities                 []message_dto.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                        `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMe1ssageId        uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}
