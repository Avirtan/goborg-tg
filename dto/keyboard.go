package dto

type Keyboard struct {
	InlineKeyboardMarkup
	ForceReply
}

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string                      `json:"text"`
	Url                          string                      `json:"url,omitempty"`
	CallbackData                 string                      `json:"callback_data,omitempty"`
	WebApp                       WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                      `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 any                         `json:"callback_game,omitempty"` // TODO 	CallbackGame
	Pay                          bool                        `json:"pay,omitempty"`
}

// https://core.telegram.org/bots/api#forcereply
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
	RequestId               int64                   `json:"request_id"`
	ChatIsChannel           bool                    `json:"chat_is_channel,omitempty"`
	ChatIsForum             bool                    `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                    `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                    `json:"chat_is_created,omitempty"`
	BotIsMember             bool                    `json:"bot_is_member,omitempty"`
	UserAdministratorRights ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
}

// https://core.telegram.org/bots/api#switchinlinequerychosenchat
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	Type int64 `json:"type,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
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

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	Id              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message,omitempty"`
	InlineMessageId string  `json:"inline_message_id,omitempty"`
	ChatInstance    string  `json:"chat_instance,omitempty"`
	Data            string  `json:"data,omitempty"`
	GameShortName   string  `json:"game_short_name,omitempty"`
}
