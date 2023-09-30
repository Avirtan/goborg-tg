package method_dto

import command_dto "TGoBot/dto/command"

// https://core.telegram.org/bots/api#setmycommands
type SetCommandRequest struct {
	Commands     []*command_dto.BotCommand `json:"commands"`
	Scope        any                       `json:"scope,omitempty"`
	LanguageCode string                    `json:"language_code,omitempty"`
}
