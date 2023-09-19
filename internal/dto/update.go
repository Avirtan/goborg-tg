package dto

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Update []Update `json:"result"`
}

type Update struct {
	UpdateID uint64   `json:"update_id"`
	Message  *Message `json:"message"`
}

type UpdateRequest struct {
	Offset  uint64 `json:"offset"`
	Limit   int    `json:"limmit"`
	Timeout int    `json:"timeout"`
}
