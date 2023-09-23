package dto

type Keyboard struct {
	InlineKeyboard
	ForceReply
}

type InlineKeyboard struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type ForceReply struct {
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	ForceReply            bool   `json:"force_reply,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}
