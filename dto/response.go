package dto

type Response struct {
	Ok          bool        `json:"ok,omitempty"`
	Result      interface{} `json:"result,omitempty"`
	Description string      `json:"description,omitempty"`
}
