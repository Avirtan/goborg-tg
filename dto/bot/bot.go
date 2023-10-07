package bot_dto

// https://core.telegram.org/bots/api#botname
type BotName struct {
	Name string `json:"name"`
}

// https://core.telegram.org/bots/api#botdescription
type BotDescription struct {
	Description string `json:"description"`
}

// https://core.telegram.org/bots/api#botshortdescription
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}
