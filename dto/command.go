package dto

// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// https://core.telegram.org/bots/api#setmycommands
type SetCommandRequest struct {
	Commands     []*BotCommand `json:"commands"`
	Scope        any           `json:"scope,omitempty"`
	LanguageCode string        `json:"language_code,omitempty"`
}
