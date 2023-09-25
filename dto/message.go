package dto

type Message struct {
	MessageId       int    `json:"message_id"`
	MessageThreadId int    `json:"message_thread_id"`
	Text            string `json:"text"`
	From            User   `json:"from"`
	SenderChat      Chat   `json:"sender_chat"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
	User          User   `json:"user,omitempty"`
}

type SendMessage struct {
	ChatID      int64    `json:"chat_id"`
	Text        string   `json:"text"`
	ReplyMarkup Keyboard `json:"reply_markup,omitempty"`
}
