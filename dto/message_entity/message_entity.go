package message_entity_dto

import user_dto "github.com/Avirtan/TGoBot/dto/user"

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          string         `json:"type"`
	Offset        int            `json:"offset"`
	Length        int            `json:"length"`
	Url           string         `json:"url,omitempty"`
	Language      string         `json:"language,omitempty"`
	CustomEmojiId string         `json:"custom_emoji_id,omitempty"`
	User          *user_dto.User `json:"user,omitempty"`
}
