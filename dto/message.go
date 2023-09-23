package dto

type Message struct {
	MessageId       int    `json:"message_id"`
	MessageThreadId int    `json:"message_thread_id"`
	Text            string `json:"text"`
	From            User   `json:"from"`
	SenderChat      Chat   `json:"sender_chat"`
}

type SendMessage struct {
	ChatID      int64    `json:"chat_id"`
	Text        string   `json:"text"`
	ReplyMarkup Keyboard `json:"reply_markup,omitempty"`
}
