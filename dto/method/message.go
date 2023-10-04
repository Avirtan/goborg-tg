package method_dto

import message_dto "github.com/Avirtan/TGoBot/dto/message"

// https://core.telegram.org/bots/api#sendmessage
type SendMessage[ID int64 | string] struct {
	ChatID                   ID                          `json:"chat_id"`
	Text                     string                      `json:"text"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	Entities                 []message_dto.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                        `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
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

// https://core.telegram.org/bots/api#copymessage
type CopyMessage struct {
	ChatID                   int64                       `json:"chat_id"`
	FromChatId               int64                       `json:"from_chat_id"`
	MessageId                int64                       `json:"message_id"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	Caption                  int64                       `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	Entities                 []message_dto.MessageEntity `json:"entities,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
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

// конструктор для SendMessage
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

func (sm *SendMessage[ID]) SetEntities(entities []message_dto.MessageEntity) *SendMessage[ID] {
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

// конструктор для ForwardMessage
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

// конструктор для CopyMessage
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

func (cm *CopyMessage) SetEntities(entities []message_dto.MessageEntity) *CopyMessage {
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
