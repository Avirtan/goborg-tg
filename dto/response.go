package dto

type Response struct {
	Ok          bool        `json:"ok,omitempty"`
	Result      interface{} `json:"result,omitempty"`
	Description string      `json:"description,omitempty"`
}

type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int64 `retry_after:"result,omitempty"`
}
