package inline_dto

import (
	user_dto "TGoBot/dto/user"
	utils_dto "TGoBot/dto/utils"
)

// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	Id       string         `json:"id"`
	From     user_dto.User  `json:"from"`
	Query    string         `json:"query"`
	Offset   string         `json:"offset"`
	Chatype  string         `json:"chat_type,omitempty"`
	Location *user_dto.User `json:"location,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultsbutton
type InlineQueryResultsButton struct {
	Text           string                `json:"text"`
	WebApp         *utils_dto.WebAppInfo `json:"web_app,omitempty"`
	StartParameter string                `json:"cache_time,omitempty"`
}

// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultId        string              `json:"result_id"`
	From            user_dto.User       `json:"from"`
	Location        *utils_dto.Location `json:"location,omitempty"`
	InlineMessageId string              `json:"inline_message_id,omitempty"`
	Query           string              `json:"query,omitempty"`
}

// https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	InlineMessageId string `json:"inline_message_id"`
}
