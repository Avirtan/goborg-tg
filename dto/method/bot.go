package method_dto

import message_dto "github.com/Avirtan/TGoBot/dto/message"

// https://core.telegram.org/bots/api#setmycommands
type GetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

// https://core.telegram.org/bots/api#setmydefaultadministratorrights
type SetMyDefaultAdministratorRights struct {
	Rights      message_dto.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                                `json:"for_channels,omitempty"`
}

// https://core.telegram.org/bots/api#getchatmenubutton
type GetChatMenuButton struct {
	ChatId uint64 `json:"chat_id,omitempty"`
}

// https://core.telegram.org/bots/api#setchatmenubutton
type SetChatMenuButton struct {
	ChatId     uint64                 `json:"chat_id,omitempty"`
	MenuButton message_dto.MenuButton `json:"menu_button,omitempty"`
}

// https://core.telegram.org/bots/api#getmyshortdescription
type GetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmyshortdescription
type SetMyShortDescription struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmydescription
type GetMyDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmydescription
type SetMyDescription struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#getmyname
type GetMyName struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#setmyname
type SetMyName struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}
