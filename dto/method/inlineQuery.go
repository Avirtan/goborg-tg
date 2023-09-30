package method_dto

import inline_dto "TGoBot/dto/inlineQuery"

// https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQuery struct {
	InlineQueryId string                               `json:"inline_query_id"`
	Results       []inline_dto.InlineQueryResult       `json:"results"`
	CacheTime     int                                  `json:"cache_time,omitempty"`
	IsPersonal    bool                                 `json:"is_personal,omitempty"`
	NextOffset    string                               `json:"next_offset,omitempty"`
	Button        *inline_dto.InlineQueryResultsButton `json:"button,omitempty"`
}

// https://core.telegram.org/bots/api#answerwebappquery
type AnswerWebAppQuery struct {
	WebAppQueryId string                       `json:"web_app_query_id"`
	Result        inline_dto.InlineQueryResult `json:"result"`
}
