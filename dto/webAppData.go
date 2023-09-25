package dto

// https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

//https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	Url string `json:"url"`
}

//https://core.telegram.org/bots/api#loginurl
type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}
