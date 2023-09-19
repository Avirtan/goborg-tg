package dto

type BotCommand struct {
	Command      string
	Description  string
	Scope        any
	LanguageCode string
}

// https://core.telegram.org/bots/api#setmycommands
type SetCommandRequest struct {
	Commands []BotCommandRequst `json:"commands"`
	// Scope        any                `json:"scope"`
	LanguageCode string `json:"language_code"`
}

// https://core.telegram.org/bots/api#botcommand
type BotCommandRequst struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}
