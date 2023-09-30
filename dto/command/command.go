package command_dto

// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}
