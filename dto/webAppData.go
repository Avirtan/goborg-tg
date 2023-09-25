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
