package method_dto

import (
	message_dto "github.com/Avirtan/goborg-tg/dto/message"
	message_entity_dto "github.com/Avirtan/goborg-tg/dto/message_entity"
)

// https://core.telegram.org/bots/api#sendmessage
type SendMessage[ID int64 | string] struct {
	ChatID                   ID                                 `json:"chat_id"`
	Text                     string                             `json:"text"`
	MessageThreadId          int                                `json:"message_thread_id,omitempty"`
	ParseMode                string                             `json:"parse_mode,omitempty"`
	Entities                 []message_entity_dto.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                               `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                               `json:"disable_notification,omitempty"`
	ProtectContent           bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewSendMessage[ID int64 | string](chatID ID, text string) *SendMessage[ID] {
	return &SendMessage[ID]{
		ChatID: chatID,
		Text:   text,
	}
}

func (sm *SendMessage[ID]) SetMessageThreadId(messageThreadId int) *SendMessage[ID] {
	sm.MessageThreadId = messageThreadId
	return sm
}

func (sm *SendMessage[ID]) SetParseMode(parseMode string) *SendMessage[ID] {
	sm.ParseMode = parseMode
	return sm
}

func (sm *SendMessage[ID]) SetEntities(entities []message_entity_dto.MessageEntity) *SendMessage[ID] {
	sm.Entities = entities
	return sm
}

func (sm *SendMessage[ID]) SetDisableWebPagePreview(disableWebPagePreview bool) *SendMessage[ID] {
	sm.DisableWebPagePreview = disableWebPagePreview
	return sm
}

func (sm *SendMessage[ID]) SetDisableNotification(disableNotification bool) *SendMessage[ID] {
	sm.DisableNotification = disableNotification
	return sm
}

func (sm *SendMessage[ID]) SetProtectContent(protectContent bool) *SendMessage[ID] {
	sm.ProtectContent = protectContent
	return sm
}

func (sm *SendMessage[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendMessage[ID] {
	sm.ReplyToMessageId = replyToMessageId
	return sm
}

func (sm *SendMessage[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendMessage[ID] {
	sm.AllowSendingWithoutReply = allowSendingWithoutReply
	return sm
}

func (sm *SendMessage[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendMessage[ID] {
	sm.ReplyMarkup = replyMarkup
	return sm
}

// https://core.telegram.org/bots/api#forwardmessage
type ForwardMessage struct {
	ChatID              int64 `json:"chat_id"`
	FromChatId          int64 `json:"from_chat_id"`
	MessageId           int64 `json:"message_id"`
	MessageThreadId     int   `json:"message_thread_id,omitempty"`
	ProtectContent      bool  `json:"protect_content,omitempty"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
}

func NewForwardMessage(chatID, fromChatID, messageID int64) *ForwardMessage {
	return &ForwardMessage{
		ChatID:     chatID,
		FromChatId: fromChatID,
		MessageId:  messageID,
	}
}

func (fm *ForwardMessage) SetMessageThreadId(messageThreadId int) *ForwardMessage {
	fm.MessageThreadId = messageThreadId
	return fm
}

func (fm *ForwardMessage) SetProtectContent(protectContent bool) *ForwardMessage {
	fm.ProtectContent = protectContent
	return fm
}

func (fm *ForwardMessage) SetDisableNotification(disableNotification bool) *ForwardMessage {
	fm.DisableNotification = disableNotification
	return fm
}

// https://core.telegram.org/bots/api#copymessage
type CopyMessage struct {
	ChatID                   int64                              `json:"chat_id"`
	FromChatId               int64                              `json:"from_chat_id"`
	MessageId                int64                              `json:"message_id"`
	MessageThreadId          int                                `json:"message_thread_id,omitempty"`
	Caption                  int64                              `json:"caption,omitempty"`
	ParseMode                string                             `json:"parse_mode,omitempty"`
	Entities                 []message_entity_dto.MessageEntity `json:"entities,omitempty"`
	DisableNotification      bool                               `json:"disable_notification,omitempty"`
	ProtectContent           bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewCopyMessage(chatID, fromChatID, messageID int64) *CopyMessage {
	return &CopyMessage{
		ChatID:     chatID,
		FromChatId: fromChatID,
		MessageId:  messageID,
	}
}

func (cm *CopyMessage) SetMessageThreadId(messageThreadId int) *CopyMessage {
	cm.MessageThreadId = messageThreadId
	return cm
}

func (cm *CopyMessage) SetCaption(caption int64) *CopyMessage {
	cm.Caption = caption
	return cm
}

func (cm *CopyMessage) SetParseMode(parseMode string) *CopyMessage {
	cm.ParseMode = parseMode
	return cm
}

func (cm *CopyMessage) SetEntities(entities []message_entity_dto.MessageEntity) *CopyMessage {
	cm.Entities = entities
	return cm
}

func (cm *CopyMessage) SetDisableNotification(disableNotification bool) *CopyMessage {
	cm.DisableNotification = disableNotification
	return cm
}

func (cm *CopyMessage) SetProtectContent(protectContent bool) *CopyMessage {
	cm.ProtectContent = protectContent
	return cm
}

func (cm *CopyMessage) SetReplyToMessageId(replyToMessageId uint64) *CopyMessage {
	cm.ReplyToMessageId = replyToMessageId
	return cm
}

func (cm *CopyMessage) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *CopyMessage {
	cm.AllowSendingWithoutReply = allowSendingWithoutReply
	return cm
}

func (cm *CopyMessage) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *CopyMessage {
	cm.ReplyMarkup = replyMarkup
	return cm
}

// https://core.telegram.org/bots/api#sendvideonote
type SendVideoNote[ID string | int64] struct {
	ChatID                   ID                    `json:"chat_id"`
	VideoNote                string                `json:"video_note"`
	MessageThreadId          int                   `json:"message_thread_id,omitempty"`
	Duration                 int                   `json:"duration,omitempty"`
	Length                   int                   `json:"length,omitempty"`
	Thumbnail                string                `json:"thumbnail,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard `json:"reply_markup,omitempty"`
}

func NewSendVideoNote[ID string | int64](chatID ID, videoNote string) *SendVideoNote[ID] {
	return &SendVideoNote[ID]{
		ChatID:    chatID,
		VideoNote: videoNote,
	}
}

func (s *SendVideoNote[ID]) SetMessageThreadId(messageThreadId int) *SendVideoNote[ID] {
	s.MessageThreadId = messageThreadId
	return s
}

func (s *SendVideoNote[ID]) SetDuration(duration int) *SendVideoNote[ID] {
	s.Duration = duration
	return s
}

func (s *SendVideoNote[ID]) SetLength(length int) *SendVideoNote[ID] {
	s.Length = length
	return s
}

func (s *SendVideoNote[ID]) SetThumbnail(thumbnail string) *SendVideoNote[ID] {
	s.Thumbnail = thumbnail
	return s
}

func (s *SendVideoNote[ID]) SetDisableNotification(disableNotification bool) *SendVideoNote[ID] {
	s.DisableNotification = disableNotification
	return s
}

func (s *SendVideoNote[ID]) SetProtectContent(protectContent bool) *SendVideoNote[ID] {
	s.ProtectContent = protectContent
	return s
}

func (s *SendVideoNote[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendVideoNote[ID] {
	s.ReplyToMessageId = replyToMessageId
	return s
}

func (s *SendVideoNote[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendVideoNote[ID] {
	s.AllowSendingWithoutReply = allowSendingWithoutReply
	return s
}

func (s *SendVideoNote[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendVideoNote[ID] {
	s.ReplyMarkup = replyMarkup
	return s
}

// https://core.telegram.org/bots/api#sendmediagroup
type SendMediaGroup[ID string | int64] struct {
	ChatID                   ID                       `json:"chat_id"`
	Media                    []message_dto.InputMedia `json:"media"`
	MessageThreadId          int                      `json:"message_thread_id,omitempty"`
	DisableNotification      bool                     `json:"disable_notification,omitempty"`
	ProtectContent           bool                     `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                     `json:"allow_sending_without_reply,omitempty"`
}

func NewSendMediaGroup[ID string | int64](chatID ID, media []message_dto.InputMedia) *SendMediaGroup[ID] {
	return &SendMediaGroup[ID]{
		ChatID: chatID,
		Media:  media,
	}
}

func (s *SendMediaGroup[ID]) SetMessageThreadId(messageThreadId int) *SendMediaGroup[ID] {
	s.MessageThreadId = messageThreadId
	return s
}

func (s *SendMediaGroup[ID]) SetDisableNotification(disableNotification bool) *SendMediaGroup[ID] {
	s.DisableNotification = disableNotification
	return s
}

func (s *SendMediaGroup[ID]) SetProtectContent(protectContent bool) *SendMediaGroup[ID] {
	s.ProtectContent = protectContent
	return s
}

func (s *SendMediaGroup[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendMediaGroup[ID] {
	s.ReplyToMessageId = replyToMessageId
	return s
}

func (s *SendMediaGroup[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendMediaGroup[ID] {
	s.AllowSendingWithoutReply = allowSendingWithoutReply
	return s
}

// https://core.telegram.org/bots/api#sendlocation
type SendLocation[ID string | int64] struct {
	ChatID                   ID                    `json:"chat_id"`
	Latitude                 float64               `json:"latitude"`
	Longitude                float64               `json:"longitude"`
	MessageThreadId          int                   `json:"message_thread_id,omitempty"`
	HorizontalAccuracy       float32               `json:"horizontal_accuracy,omitempty"`
	LivePeriod               int                   `json:"live_period,omitempty"`
	Heading                  uint16                `json:"heading,omitempty"`
	ProximityAlertRadius     uint                  `json:"proximity_alert_radius,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard `json:"reply_markup,omitempty"`
}

func NewSendLocation[ID string | int64](chatID ID, latitude float64, longitude float64) *SendLocation[ID] {
	return &SendLocation[ID]{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (s *SendLocation[ID]) SetMessageThreadId(messageThreadId int) *SendLocation[ID] {
	s.MessageThreadId = messageThreadId
	return s
}

func (s *SendLocation[ID]) SetHorizontalAccuracy(horizontalAccuracy float32) *SendLocation[ID] {
	s.HorizontalAccuracy = horizontalAccuracy
	return s
}

func (s *SendLocation[ID]) SetLivePeriod(livePeriod int) *SendLocation[ID] {
	s.LivePeriod = livePeriod
	return s
}

func (s *SendLocation[ID]) SetHeading(heading uint16) *SendLocation[ID] {
	s.Heading = heading
	return s
}

func (s *SendLocation[ID]) SetProximityAlertRadius(proximityAlertRadius uint) *SendLocation[ID] {
	s.ProximityAlertRadius = proximityAlertRadius
	return s
}

func (s *SendLocation[ID]) SetDisableNotification(disableNotification bool) *SendLocation[ID] {
	s.DisableNotification = disableNotification
	return s
}

func (s *SendLocation[ID]) SetProtectContent(protectContent bool) *SendLocation[ID] {
	s.ProtectContent = protectContent
	return s
}

func (s *SendLocation[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendLocation[ID] {
	s.ReplyToMessageId = replyToMessageId
	return s
}

func (s *SendLocation[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendLocation[ID] {
	s.AllowSendingWithoutReply = allowSendingWithoutReply
	return s
}

// https://core.telegram.org/bots/api#editmessagetext
type EditMessageText[ID int64 | string] struct {
	ChatID                ID                                 `json:"chat_id"`
	MessageID             int                                `json:"message_id,omitempty"`
	InlineMessageID       string                             `json:"inline_message_id,omitempty"`
	Text                  string                             `json:"text"`
	ParseMode             string                             `json:"parse_mode,omitempty"`
	Entities              []message_entity_dto.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                               `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewEditMessageText[ID int64 | string](chatID ID, text string) *EditMessageText[ID] {
	return &EditMessageText[ID]{
		ChatID: chatID,
		Text:   text,
	}
}

func (s *EditMessageText[ID]) SetMessageID(messageID int) *EditMessageText[ID] {
	s.MessageID = messageID
	return s
}

func (s *EditMessageText[ID]) SetInlineMessageID(inlineMessageID string) *EditMessageText[ID] {
	s.InlineMessageID = inlineMessageID
	return s
}

func (s *EditMessageText[ID]) SetParseMode(parseMode string) *EditMessageText[ID] {
	s.ParseMode = parseMode
	return s
}

func (s *EditMessageText[ID]) SetEntities(entities []message_entity_dto.MessageEntity) *EditMessageText[ID] {
	s.Entities = entities
	return s
}

func (s *EditMessageText[ID]) SetDisableWebPagePreview(disableWebPagePreview bool) *EditMessageText[ID] {
	s.DisableWebPagePreview = disableWebPagePreview
	return s
}

func (s *EditMessageText[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *EditMessageText[ID] {
	s.ReplyMarkup = replyMarkup
	return s
}

// https://core.telegram.org/bots/api#editmessagecaption
type EditMessageCaption[ID int64 | string] struct {
	ChatID          ID                                 `json:"chat_id,omitempty"`
	MessageID       int                                `json:"message_id,omitempty"`
	InlineMessageID string                             `json:"inline_message_id,omitempty"`
	Caption         string                             `json:"caption,omitempty"`
	ParseMode       string                             `json:"parse_mode,omitempty"`
	CaptionEntities []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	ReplyMarkup     message_dto.InlineKeyboardMarkup   `json:"reply_markup,omitempty"`
}

func NewEditMessageCaption[ID int64 | string](text string) *EditMessageCaption[ID] {
	return &EditMessageCaption[ID]{
		Caption: text,
	}
}

func (e *EditMessageCaption[ID]) SetChatID(chatID ID) *EditMessageCaption[ID] {
	e.ChatID = chatID
	return e
}

func (e *EditMessageCaption[ID]) SetMessageID(messageID int) *EditMessageCaption[ID] {
	e.MessageID = messageID
	return e
}

func (e *EditMessageCaption[ID]) SetInlineMessageID(inlineMessageID string) *EditMessageCaption[ID] {
	e.InlineMessageID = inlineMessageID
	return e
}

func (e *EditMessageCaption[ID]) SetParseMode(parseMode string) *EditMessageCaption[ID] {
	e.ParseMode = parseMode
	return e
}

func (e *EditMessageCaption[ID]) SetCaptionEntities(entities []message_entity_dto.MessageEntity) *EditMessageCaption[ID] {
	e.CaptionEntities = entities
	return e
}

func (e *EditMessageCaption[ID]) SetReplyMarkup(replyMarkup message_dto.InlineKeyboardMarkup) *EditMessageCaption[ID] {
	e.ReplyMarkup = replyMarkup
	return e
}

// https://core.telegram.org/bots/api#editmessagemedia
type EditMessageMedia[ID int64 | string] struct {
	ChatID          ID                               `json:"chat_id,omitempty"`
	MessageID       int                              `json:"message_id,omitempty"`
	InlineMessageID string                           `json:"inline_message_id,omitempty"`
	Media           message_dto.InputMedia           `json:"media"`
	ReplyMarkup     message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewEditMessageMedia[ID int64 | string](media message_dto.InputMedia) *EditMessageMedia[ID] {
	return &EditMessageMedia[ID]{
		Media: media,
	}
}

func (e *EditMessageMedia[ID]) SetChatID(chatID ID) *EditMessageMedia[ID] {
	e.ChatID = chatID
	return e
}

func (e *EditMessageMedia[ID]) SetMessageID(messageID int) *EditMessageMedia[ID] {
	e.MessageID = messageID
	return e
}

func (e *EditMessageMedia[ID]) SetInlineMessageID(inlineMessageID string) *EditMessageMedia[ID] {
	e.InlineMessageID = inlineMessageID
	return e
}

func (e *EditMessageMedia[ID]) SetReplyMarkup(replyMarkup message_dto.InlineKeyboardMarkup) *EditMessageMedia[ID] {
	e.ReplyMarkup = replyMarkup
	return e
}

// https://core.telegram.org/bots/api#editmessagelivelocation
type EditMessageLiveLocation[ID int64 | string] struct {
	ChatID               ID                               `json:"chat_id,omitempty"`
	MessageID            int                              `json:"message_id,omitempty"`
	InlineMessageID      string                           `json:"inline_message_id,omitempty"`
	Latitude             float64                          `json:"latitude"`
	Longitude            float64                          `json:"longitude"`
	HorizontalAccuracy   float64                          `json:"horizontal_accuracy,omitempty"`
	Heading              int                              `json:"heading,omitempty"`
	ProximityAlertRadius int                              `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewEditMessageLiveLocation[ID int64 | string](latitude, longitude float64) *EditMessageLiveLocation[ID] {
	return &EditMessageLiveLocation[ID]{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (e *EditMessageLiveLocation[ID]) SetChatID(chatID ID) *EditMessageLiveLocation[ID] {
	e.ChatID = chatID
	return e
}

func (e *EditMessageLiveLocation[ID]) SetMessageID(messageID int) *EditMessageLiveLocation[ID] {
	e.MessageID = messageID
	return e
}

func (e *EditMessageLiveLocation[ID]) SetInlineMessageID(inlineMessageID string) *EditMessageLiveLocation[ID] {
	e.InlineMessageID = inlineMessageID
	return e
}

func (e *EditMessageLiveLocation[ID]) SetHorizontalAccuracy(horizontalAccuracy float64) *EditMessageLiveLocation[ID] {
	e.HorizontalAccuracy = horizontalAccuracy
	return e
}

func (e *EditMessageLiveLocation[ID]) SetHeading(heading int) *EditMessageLiveLocation[ID] {
	e.Heading = heading
	return e
}

func (e *EditMessageLiveLocation[ID]) SetProximityAlertRadius(proximityAlertRadius int) *EditMessageLiveLocation[ID] {
	e.ProximityAlertRadius = proximityAlertRadius
	return e
}

func (e *EditMessageLiveLocation[ID]) SetReplyMarkup(replyMarkup message_dto.InlineKeyboardMarkup) *EditMessageLiveLocation[ID] {
	e.ReplyMarkup = replyMarkup
	return e
}

// https://core.telegram.org/bots/api#stopmessagelivelocation
type StopMessageLiveLocation[ID int64 | string] struct {
	ChatID          ID                               `json:"chat_id,omitempty"`
	MessageID       int                              `json:"message_id,omitempty"`
	InlineMessageID string                           `json:"inline_message_id,omitempty"`
	ReplyMarkup     message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewStopMessageLiveLocation[ID int64 | string]() *StopMessageLiveLocation[ID] {
	return &StopMessageLiveLocation[ID]{}
}

func (s *StopMessageLiveLocation[ID]) SetChatID(chatID ID) *StopMessageLiveLocation[ID] {
	s.ChatID = chatID
	return s
}

func (s *StopMessageLiveLocation[ID]) SetMessageID(messageID int) *StopMessageLiveLocation[ID] {
	s.MessageID = messageID
	return s
}

func (s *StopMessageLiveLocation[ID]) SetInlineMessageID(inlineMessageID string) *StopMessageLiveLocation[ID] {
	s.InlineMessageID = inlineMessageID
	return s
}

func (s *StopMessageLiveLocation[ID]) SetReplyMarkup(replyMarkup message_dto.InlineKeyboardMarkup) *StopMessageLiveLocation[ID] {
	s.ReplyMarkup = replyMarkup
	return s
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkup[ID int64 | string] struct {
	ChatID          ID                               `json:"chat_id,omitempty"`
	MessageID       int                              `json:"message_id,omitempty"`
	InlineMessageID string                           `json:"inline_message_id,omitempty"`
	ReplyMarkup     message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewEditMessageReplyMarkup[ID int64 | string]() *EditMessageReplyMarkup[ID] {
	return &EditMessageReplyMarkup[ID]{}
}

func (e *EditMessageReplyMarkup[ID]) SetChatID(chatID ID) *EditMessageReplyMarkup[ID] {
	e.ChatID = chatID
	return e
}

func (e *EditMessageReplyMarkup[ID]) SetMessageID(messageID int) *EditMessageReplyMarkup[ID] {
	e.MessageID = messageID
	return e
}

func (e *EditMessageReplyMarkup[ID]) SetInlineMessageID(inlineMessageID string) *EditMessageReplyMarkup[ID] {
	e.InlineMessageID = inlineMessageID
	return e
}

func (e *EditMessageReplyMarkup[ID]) SetReplyMarkup(replyMarkup message_dto.InlineKeyboardMarkup) *EditMessageReplyMarkup[ID] {
	e.ReplyMarkup = replyMarkup
	return e
}

// https://core.telegram.org/bots/api#stoppoll
type StopPoll[ID int64 | string] struct {
	ChatID      ID                               `json:"chat_id"`
	MessageID   int                              `json:"message_id"`
	ReplyMarkup message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewStopPoll[ID int64 | string](chatID ID, messageID int) *StopPoll[ID] {
	return &StopPoll[ID]{
		ChatID:    chatID,
		MessageID: messageID,
	}
}

func (s *StopPoll[ID]) SetReplyMarkup(replyMarkup message_dto.InlineKeyboardMarkup) *StopPoll[ID] {
	s.ReplyMarkup = replyMarkup
	return s
}

// https://core.telegram.org/bots/api#deletemessage
type DeleteMessage[ID int64 | string] struct {
	ChatID    ID  `json:"chat_id"`
	MessageID int `json:"message_id"`
}

func NewDeleteMessage[ID int64 | string](chatID ID, messageID int) *DeleteMessage[ID] {
	return &DeleteMessage[ID]{
		ChatID:    chatID,
		MessageID: messageID,
	}
}
