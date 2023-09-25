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

// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string                    `json:"text"`
	RequestUser     KeyboardButtonRequestUser `json:"is_persistent,omitempty"`
	RequestChat     KeyboardButtonRequestChat `json:"request_chat,omitempty"`
	RequestContact  bool                      `json:"request_contact,omitempty"`
	RequestLocation bool                      `json:"request_location,omitempty"`
	RequestPoll     KeyboardButtonPollType    `json:"request_poll,omitempty"`
	WebApp          WebAppInfo                `json:"web_app,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbuttonrequestuser
type KeyboardButtonRequestUser struct {
	RequestId     int64 `json:"request_id"`
	UserIsNot     bool  `json:"user_is_bot,omitempty"`
	UserIsPremium bool  `json:"user_is_premium,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbuttonrequestchat
type KeyboardButtonRequestChat struct {
	RequestId               int64 `json:"request_id"`
	ChatIsChannel           bool  `json:"chat_is_channel,omitempty"`
	ChatIsForum             bool  `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool  `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool  `json:"chat_is_created,omitempty"`
	BotIsMember             bool  `json:"bot_is_member,omitempty"`
	UserAdministratorRights any   `json:"user_administrator_rights,omitempty"` // TODO ChatAdministratorRights
	BotAdministratorRights  any   `json:"bot_administrator_rights,omitempty"`  // TODO ChatAdministratorRights
}

// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	Type int64 `json:"type,omitempty"`
}

// https://core.telegram.org/bots/api#usershared
type UserShared struct {
	RequestId int64 `json:"request_id"`
	UserId    int64 `json:"user_id"`
}

// https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	RequestId int64 `json:"request_id"`
	СhatId    int64 `json:"chat_id"`
}
