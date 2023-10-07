package method_dto

import inline_dto "github.com/Avirtan/TGoBot/dto/inlineQuery"

// https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQuery struct {
	InlineQueryId string                               `json:"inline_query_id"`
	Results       []inline_dto.InlineQueryResult       `json:"results"`
	CacheTime     int                                  `json:"cache_time,omitempty"`
	IsPersonal    bool                                 `json:"is_personal,omitempty"`
	NextOffset    string                               `json:"next_offset,omitempty"`
	Button        *inline_dto.InlineQueryResultsButton `json:"button,omitempty"`
}

func NewAnswerInlineQuery(inlineQueryId string, results []inline_dto.InlineQueryResult) *AnswerInlineQuery {
	return &AnswerInlineQuery{
		InlineQueryId: inlineQueryId,
		Results:       results,
	}
}

func (aiq *AnswerInlineQuery) SetCacheTime(cacheTime int) *AnswerInlineQuery {
	aiq.CacheTime = cacheTime
	return aiq
}

func (aiq *AnswerInlineQuery) SetIsPersonal(isPersonal bool) *AnswerInlineQuery {
	aiq.IsPersonal = isPersonal
	return aiq
}

func (aiq *AnswerInlineQuery) SetNextOffset(nextOffset string) *AnswerInlineQuery {
	aiq.NextOffset = nextOffset
	return aiq
}

func (aiq *AnswerInlineQuery) SetButton(button *inline_dto.InlineQueryResultsButton) *AnswerInlineQuery {
	aiq.Button = button
	return aiq
}

// https://core.telegram.org/bots/api#answerwebappquery
type AnswerWebAppQuery struct {
	WebAppQueryId string                       `json:"web_app_query_id"`
	Result        inline_dto.InlineQueryResult `json:"result"`
}

func NewAnswerWebAppQuery(webAppQueryId string, result inline_dto.InlineQueryResult) *AnswerWebAppQuery {
	return &AnswerWebAppQuery{
		WebAppQueryId: webAppQueryId,
		Result:        result,
	}
}

// https://core.telegram.org/bots/api#answercallbackquery
type AnswerCallbackQuery struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       uint64 `json:"cache_time,omitempty"`
}

func NewAnswerCallbackQuery(callbackQueryId string) *AnswerCallbackQuery {
	return &AnswerCallbackQuery{
		CallbackQueryId: callbackQueryId,
	}
}

func (a *AnswerCallbackQuery) SetText(text string) *AnswerCallbackQuery {
	a.Text = text
	return a
}

func (a *AnswerCallbackQuery) SetShowAlert(showAlert bool) *AnswerCallbackQuery {
	a.ShowAlert = showAlert
	return a
}

func (a *AnswerCallbackQuery) SetUrl(url string) *AnswerCallbackQuery {
	a.Url = url
	return a
}

func (a *AnswerCallbackQuery) SetCacheTime(cacheTime uint64) *AnswerCallbackQuery {
	a.CacheTime = cacheTime
	return a
}
