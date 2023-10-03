package method_dto

import command_dto "TGoBot/dto/command"

// https://core.telegram.org/bots/api#setmycommands
type SetCommandRequest struct {
	Commands     []*command_dto.BotCommand `json:"commands"`
	Scope        any                       `json:"scope,omitempty"`
	LanguageCode string                    `json:"language_code,omitempty"`
}

func NewSetCommandRequest(commands []*command_dto.BotCommand) *SetCommandRequest {
	return &SetCommandRequest{
		Commands: commands,
	}
}

func (scr *SetCommandRequest) SetScope(scope any) *SetCommandRequest {
	scr.Scope = scope
	return scr
}

func (scr *SetCommandRequest) SetLanguageCode(languageCode string) *SetCommandRequest {
	scr.LanguageCode = languageCode
	return scr
}
