package method_dto

import message_dto "github.com/Avirtan/goborg-tg/dto/message"

// https://core.telegram.org/bots/api#setmycommands
type GetMyDefaultAdministratorRights struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

func NewGetMyDefaultAdministratorRights() *GetMyDefaultAdministratorRights {
	return &GetMyDefaultAdministratorRights{}
}

func (g *GetMyDefaultAdministratorRights) SetForChannels(forChannels bool) *GetMyDefaultAdministratorRights {
	g.ForChannels = forChannels
	return g
}

// https://core.telegram.org/bots/api#setmydefaultadministratorrights
type SetMyDefaultAdministratorRights struct {
	Rights      message_dto.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                                `json:"for_channels,omitempty"`
}

func NewSetMyDefaultAdministratorRights() *SetMyDefaultAdministratorRights {
	return &SetMyDefaultAdministratorRights{}
}

func (s *SetMyDefaultAdministratorRights) SetRights(rights message_dto.ChatAdministratorRights) *SetMyDefaultAdministratorRights {
	s.Rights = rights
	return s
}

func (s *SetMyDefaultAdministratorRights) SetForChannels(forChannels bool) *SetMyDefaultAdministratorRights {
	s.ForChannels = forChannels
	return s
}

// https://core.telegram.org/bots/api#getchatmenubutton
type GetChatMenuButton struct {
	ChatId uint64 `json:"chat_id,omitempty"`
}

func NewGetChatMenuButton() *GetChatMenuButton {
	return &GetChatMenuButton{}
}

func (g *GetChatMenuButton) SetChatId(chatId uint64) *GetChatMenuButton {
	g.ChatId = chatId
	return g
}

// https://core.telegram.org/bots/api#setchatmenubutton
type SetChatMenuButton struct {
	ChatId     uint64                 `json:"chat_id,omitempty"`
	MenuButton message_dto.MenuButton `json:"menu_button,omitempty"`
}

func NewSetChatMenuButton() *SetChatMenuButton {
	return &SetChatMenuButton{}
}

func (s *SetChatMenuButton) SetChatId(chatId uint64) *SetChatMenuButton {
	s.ChatId = chatId
	return s
}

func (s *SetChatMenuButton) SetMenuButton(menuButton message_dto.MenuButton) *SetChatMenuButton {
	s.MenuButton = menuButton
	return s
}

// https://core.telegram.org/bots/api#getmyshortdescription
type GetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

func NewGetMyShortDescription() *GetMyShortDescription {
	return &GetMyShortDescription{}
}

func (g *GetMyShortDescription) SetLanguageCode(languageCode string) *GetMyShortDescription {
	g.LanguageCode = languageCode
	return g
}

// https://core.telegram.org/bots/api#setmyshortdescription
type SetMyShortDescription struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

func NewSetMyShortDescription() *SetMyShortDescription {
	return &SetMyShortDescription{}
}

func (s *SetMyShortDescription) SetShortDescription(shortDescription string) *SetMyShortDescription {
	s.ShortDescription = shortDescription
	return s
}

func (s *SetMyShortDescription) SetLanguageCode(languageCode string) *SetMyShortDescription {
	s.LanguageCode = languageCode
	return s
}

// https://core.telegram.org/bots/api#getmydescription
type GetMyDescription struct {
	LanguageCode string `json:"language_code,omitempty"`
}

func NewGetMyDescription() *GetMyDescription {
	return &GetMyDescription{}
}

func (g *GetMyDescription) SetLanguageCode(languageCode string) *GetMyDescription {
	g.LanguageCode = languageCode
	return g
}

// https://core.telegram.org/bots/api#setmydescription
type SetMyDescription struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

func NewSetMyDescription() *SetMyDescription {
	return &SetMyDescription{}
}

func (s *SetMyDescription) SetDescription(description string) *SetMyDescription {
	s.Description = description
	return s
}

func (s *SetMyDescription) SetLanguageCode(languageCode string) *SetMyDescription {
	s.LanguageCode = languageCode
	return s
}

// https://core.telegram.org/bots/api#getmyname
type GetMyName struct {
	LanguageCode string `json:"language_code,omitempty"`
}

func NewGetMyName() *GetMyName {
	return &GetMyName{}
}

func (g *GetMyName) SetLanguageCode(languageCode string) *GetMyName {
	g.LanguageCode = languageCode
	return g
}

// https://core.telegram.org/bots/api#setmyname
type SetMyName struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

func NewSetMyName() *SetMyName {
	return &SetMyName{}
}

func (s *SetMyName) SetName(name string) *SetMyName {
	s.Name = name
	return s
}

func (s *SetMyName) SetLanguageCode(languageCode string) *SetMyName {
	s.LanguageCode = languageCode
	return s
}

// https://core.telegram.org/bots/api#getfile
type GetFile struct {
	FileId string `json:"file_id,omitempty"`
}

func NewGetFile() *GetFile {
	return &GetFile{}
}

func (g *GetFile) SetFileId(fileId string) *GetFile {
	g.FileId = fileId
	return g
}
