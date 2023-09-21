package dto

type Message struct {
	MessageId       int    `json:"message_id"`
	MessageThreadId int    `json:"message_thread_id"`
	From            User   `json:"from"`
	SenderChat      Chat   `json:"sender_chat"`
	Text            string `json:"text"`
}

type SendMessage struct {
	ChatID     int        `json:"chat_id"`
	Text       string     `json:"text"`
	ForceReply ForceReply `json:"reply_markup,omitempty"`
}
