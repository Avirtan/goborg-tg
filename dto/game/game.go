package game_dto

import (
	message_dto "github.com/Avirtan/TGoBot/dto/message"
	user_dto "github.com/Avirtan/TGoBot/dto/user"
	utils_dto "github.com/Avirtan/TGoBot/dto/utils"
)

// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string                      `json:"title"`
	Description  string                      `json:"description"`
	Photo        []utils_dto.PhotoSize       `json:"photo"`
	Text         string                      `json:"text,omitempty"`
	TextEntities []message_dto.MessageEntity `json:"text_entities,omitempty"`
	Animation    utils_dto.Animation         `json:"animation,omitempty"`
}

// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int           `json:"position"`
	User     user_dto.User `json:"user"`
	Score    int           `json:"score"`
}
