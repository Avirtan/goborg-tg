package dto

// https://core.telegram.org/bots/api#botcommandscope
type BotCommandScope struct {
	Type   string `json:"type"`
	ChatId int64  `json:"chat_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}
