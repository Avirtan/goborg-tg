package method_dto

import message_dto "TGoBot/dto/message"

// https://core.telegram.org/bots/api#sendmessage
type SendMessage[ID int64 | string] struct {
	ChatID                   ID                          `json:"chat_id"`
	Text                     string                      `json:"text"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	Entities                 []message_dto.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                        `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#forwardmessage
type ForwardMessage struct {
	ChatID              int64 `json:"chat_id"`
	FromChatId          int64 `json:"from_chat_id"`
	MessageId           int64 `json:"message_id"`
	MessageThreadId     int   `json:"message_thread_id,omitempty"`
	ProtectContent      bool  `json:"protect_content,omitempty"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
}

// https://core.telegram.org/bots/api#copymessage
type CopyMessage struct {
	ChatID                   int64                       `json:"chat_id"`
	FromChatId               int64                       `json:"from_chat_id"`
	MessageId                int64                       `json:"message_id"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	Caption                  int64                       `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	Entities                 []message_dto.MessageEntity `json:"entities,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}
