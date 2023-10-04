package method_dto

import message_dto "github.com/Avirtan/TGoBot/dto/message"

// https://core.telegram.org/bots/api#sendgame
type SendGame struct {
	ChatID                   int64                             `json:"chat_id"`
	GameShortName            string                            `json:"game_short_name"`
	MessageThreadId          int                               `json:"message_thread_id,omitempty"`
	DisableNotification      bool                              `json:"disable_notification,omitempty"`
	ProtectContent           bool                              `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                              `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewSendGame(chatID int64, gameShortName string) *SendGame {
	return &SendGame{
		ChatID:        chatID,
		GameShortName: gameShortName,
	}
}

func (sg *SendGame) SetMessageThreadId(messageThreadId int) *SendGame {
	sg.MessageThreadId = messageThreadId
	return sg
}

func (sg *SendGame) SetDisableNotification(disable bool) *SendGame {
	sg.DisableNotification = disable
	return sg
}

func (sg *SendGame) SetProtectContent(protect bool) *SendGame {
	sg.ProtectContent = protect
	return sg
}

func (sg *SendGame) SetReplyToMessageId(messageID uint64) *SendGame {
	sg.ReplyToMessageId = messageID
	return sg
}

func (sg *SendGame) SetAllowSendingWithoutReply(allow bool) *SendGame {
	sg.AllowSendingWithoutReply = allow
	return sg
}

func (sg *SendGame) SetReplyMarkup(markup *message_dto.InlineKeyboardMarkup) *SendGame {
	sg.ReplyMarkup = markup
	return sg
}

// https://core.telegram.org/bots/api#setgamescore
type SetGameScore struct {
	UserID             int64  `json:"user_id"`
	Score              int    `json:"score"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	ChatID             int64  `json:"chat_id,omitempty"`
	MessageID          int    `json:"message_id,omitempty"`
	InlineMessageID    string `json:"inline_message_id,omitempty"`
}

func NewSetGameScore(userID int64, score int) *SetGameScore {
	return &SetGameScore{
		UserID: userID,
		Score:  score,
	}
}

func (sgs *SetGameScore) SetForce(force bool) *SetGameScore {
	sgs.Force = force
	return sgs
}

func (sgs *SetGameScore) SetDisableEditMessage(disable bool) *SetGameScore {
	sgs.DisableEditMessage = disable
	return sgs
}

func (sgs *SetGameScore) SetChatID(chatID int64) *SetGameScore {
	sgs.ChatID = chatID
	return sgs
}

func (sgs *SetGameScore) SetMessageID(messageID int) *SetGameScore {
	sgs.MessageID = messageID
	return sgs
}

func (sgs *SetGameScore) SetInlineMessageID(inlineMessageID string) *SetGameScore {
	sgs.InlineMessageID = inlineMessageID
	return sgs
}

// https://core.telegram.org/bots/api#getgamehighscores
type GetGameHighScores struct {
	UserID          int64  `json:"user_id"`
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

func NewGetGameHighScores(userID int64) *GetGameHighScores {
	return &GetGameHighScores{
		UserID: userID,
	}
}

func (gghs *GetGameHighScores) SetChatID(chatID int64) *GetGameHighScores {
	gghs.ChatID = chatID
	return gghs
}

func (gghs *GetGameHighScores) SetMessageID(messageID int) *GetGameHighScores {
	gghs.MessageID = messageID
	return gghs
}

func (gghs *GetGameHighScores) SetInlineMessageID(inlineMessageID string) *GetGameHighScores {
	gghs.InlineMessageID = inlineMessageID
	return gghs
}
