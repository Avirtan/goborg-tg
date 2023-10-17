package method_dto

import (
	message_dto "github.com/Avirtan/goborg-tg/dto/message"
	message_entity_dto "github.com/Avirtan/goborg-tg/dto/message_entity"
)

// https://core.telegram.org/bots/api#sendphoto
type SendPhoto[ID string | int64] struct {
	ChatID                   ID                                 `json:"chat_id"`
	Photo                    string                             `json:"photo"`
	MessageThreadId          int                                `json:"message_thread_id,omitempty"`
	Caption                  string                             `json:"caption,omitempty"`
	ParseMode                string                             `json:"parse_mode,omitempty"`
	CaptionEntities          []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                               `json:"has_spoiler,omitempty"`
	DisableNotification      bool                               `json:"disable_notification,omitempty"`
	ProtectContent           bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewSendPhoto[ID string | int64](chatID ID, photo string) *SendPhoto[ID] {
	return &SendPhoto[ID]{
		ChatID: chatID,
		Photo:  photo,
	}
}

func (sp *SendPhoto[ID]) SetMessageThreadId(messageThreadId int) *SendPhoto[ID] {
	sp.MessageThreadId = messageThreadId
	return sp
}

func (sp *SendPhoto[ID]) SetCaption(caption string) *SendPhoto[ID] {
	sp.Caption = caption
	return sp
}

func (sp *SendPhoto[ID]) SetParseMode(parseMode string) *SendPhoto[ID] {
	sp.ParseMode = parseMode
	return sp
}

func (sp *SendPhoto[ID]) SetCaptionEntities(captionEntities []message_entity_dto.MessageEntity) *SendPhoto[ID] {
	sp.CaptionEntities = captionEntities
	return sp
}

func (sp *SendPhoto[ID]) SetHasSpoiler(hasSpoiler bool) *SendPhoto[ID] {
	sp.HasSpoiler = hasSpoiler
	return sp
}

func (sp *SendPhoto[ID]) SetDisableNotification(disableNotification bool) *SendPhoto[ID] {
	sp.DisableNotification = disableNotification
	return sp
}

func (sp *SendPhoto[ID]) SetProtectContent(protectContent bool) *SendPhoto[ID] {
	sp.ProtectContent = protectContent
	return sp
}

func (sp *SendPhoto[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendPhoto[ID] {
	sp.ReplyToMessageId = replyToMessageId
	return sp
}

func (sp *SendPhoto[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendPhoto[ID] {
	sp.AllowSendingWithoutReply = allowSendingWithoutReply
	return sp
}

func (sp *SendPhoto[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendPhoto[ID] {
	sp.ReplyMarkup = replyMarkup
	return sp
}

// https://core.telegram.org/bots/api#sendaudio
type SendAudio[ID string | int64] struct {
	ChatID                   ID                                 `json:"chat_id"`
	Audio                    string                             `json:"audio"`
	MessageThreadId          int                                `json:"message_thread_id,omitempty"`
	Caption                  string                             `json:"caption,omitempty"`
	ParseMode                string                             `json:"parse_mode,omitempty"`
	CaptionEntities          []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                                `json:"duration,omitempty"`
	Performer                string                             `json:"performer,omitempty"`
	Title                    string                             `json:"title,omitempty"`
	Thumbnail                string                             `json:"thumbnail,omitempty"`
	DisableNotification      bool                               `json:"disable_notification,omitempty"`
	ProtectContent           bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewSendAudio[ID string | int64](chatID ID, audio string) *SendAudio[ID] {
	return &SendAudio[ID]{
		ChatID: chatID,
		Audio:  audio,
	}
}

func (sa *SendAudio[ID]) SetMessageThreadId(messageThreadId int) *SendAudio[ID] {
	sa.MessageThreadId = messageThreadId
	return sa
}

func (sa *SendAudio[ID]) SetCaption(caption string) *SendAudio[ID] {
	sa.Caption = caption
	return sa
}

func (sa *SendAudio[ID]) SetParseMode(parseMode string) *SendAudio[ID] {
	sa.ParseMode = parseMode
	return sa
}

func (sa *SendAudio[ID]) SetCaptionEntities(captionEntities []message_entity_dto.MessageEntity) *SendAudio[ID] {
	sa.CaptionEntities = captionEntities
	return sa
}

func (sa *SendAudio[ID]) SetDuration(duration int) *SendAudio[ID] {
	sa.Duration = duration
	return sa
}

func (sa *SendAudio[ID]) SetPerformer(performer string) *SendAudio[ID] {
	sa.Performer = performer
	return sa
}

func (sa *SendAudio[ID]) SetTitle(title string) *SendAudio[ID] {
	sa.Title = title
	return sa
}

func (sa *SendAudio[ID]) SetThumbnail(thumbnail string) *SendAudio[ID] {
	sa.Thumbnail = thumbnail
	return sa
}

func (sa *SendAudio[ID]) SetDisableNotification(disableNotification bool) *SendAudio[ID] {
	sa.DisableNotification = disableNotification
	return sa
}

func (sa *SendAudio[ID]) SetProtectContent(protectContent bool) *SendAudio[ID] {
	sa.ProtectContent = protectContent
	return sa
}

func (sa *SendAudio[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendAudio[ID] {
	sa.ReplyToMessageId = replyToMessageId
	return sa
}

func (sa *SendAudio[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendAudio[ID] {
	sa.AllowSendingWithoutReply = allowSendingWithoutReply
	return sa
}

func (sa *SendAudio[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendAudio[ID] {
	sa.ReplyMarkup = replyMarkup
	return sa
}

// https://core.telegram.org/bots/api#senddocument
type SendDocument[ID string | int64] struct {
	ChatID                      ID                                 `json:"chat_id"`
	Document                    string                             `json:"document"`
	Thumbnail                   string                             `json:"thumbnail,omitempty"`
	MessageThreadId             int                                `json:"message_thread_id,omitempty"`
	Caption                     string                             `json:"caption,omitempty"`
	ParseMode                   string                             `json:"parse_mode,omitempty"`
	CaptionEntities             []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                               `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                               `json:"disable_notification,omitempty"`
	ProtectContent              bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId            uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewSendDocument[ID string | int64](chatID ID, document string) *SendDocument[ID] {
	return &SendDocument[ID]{
		ChatID:   chatID,
		Document: document,
	}
}

func (sd *SendDocument[ID]) SetThumbnail(thumbnail string) *SendDocument[ID] {
	sd.Thumbnail = thumbnail
	return sd
}

func (sd *SendDocument[ID]) SetMessageThreadId(messageThreadId int) *SendDocument[ID] {
	sd.MessageThreadId = messageThreadId
	return sd
}

func (sd *SendDocument[ID]) SetCaption(caption string) *SendDocument[ID] {
	sd.Caption = caption
	return sd
}

func (sd *SendDocument[ID]) SetParseMode(parseMode string) *SendDocument[ID] {
	sd.ParseMode = parseMode
	return sd
}

func (sd *SendDocument[ID]) SetCaptionEntities(captionEntities []message_entity_dto.MessageEntity) *SendDocument[ID] {
	sd.CaptionEntities = captionEntities
	return sd
}

func (sd *SendDocument[ID]) SetDisableContentTypeDetection(disableContentTypeDetection bool) *SendDocument[ID] {
	sd.DisableContentTypeDetection = disableContentTypeDetection
	return sd
}

func (sd *SendDocument[ID]) SetDisableNotification(disableNotification bool) *SendDocument[ID] {
	sd.DisableNotification = disableNotification
	return sd
}

func (sd *SendDocument[ID]) SetProtectContent(protectContent bool) *SendDocument[ID] {
	sd.ProtectContent = protectContent
	return sd
}

func (sd *SendDocument[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendDocument[ID] {
	sd.ReplyToMessageId = replyToMessageId
	return sd
}

func (sd *SendDocument[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendDocument[ID] {
	sd.AllowSendingWithoutReply = allowSendingWithoutReply
	return sd
}

func (sd *SendDocument[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendDocument[ID] {
	sd.ReplyMarkup = replyMarkup
	return sd
}

// https://core.telegram.org/bots/api#sendvideo
type SendVideo[ID string | int64] struct {
	ChatID                   ID                                 `json:"chat_id"`
	Video                    string                             `json:"video"`
	MessageThreadId          int                                `json:"message_thread_id,omitempty"`
	Duration                 int                                `json:"duration,omitempty"`
	Width                    int                                `json:"width,omitempty"`
	Height                   int                                `json:"height,omitempty"`
	Thumbnail                string                             `json:"thumbnail,omitempty"`
	Caption                  string                             `json:"caption,omitempty"`
	ParseMode                string                             `json:"parse_mode,omitempty"`
	CaptionEntities          []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                               `json:"has_spoiler,omitempty"`
	SupportsStreaming        bool                               `json:"supports_streaming,omitempty"`
	DisableNotification      bool                               `json:"disable_notification,omitempty"`
	ProtectContent           bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewSendVideo[ID string | int64](chatID ID, video string) *SendVideo[ID] {
	return &SendVideo[ID]{
		ChatID: chatID,
		Video:  video,
	}
}

func (sv *SendVideo[ID]) SetMessageThreadId(messageThreadId int) *SendVideo[ID] {
	sv.MessageThreadId = messageThreadId
	return sv
}

func (sv *SendVideo[ID]) SetDuration(duration int) *SendVideo[ID] {
	sv.Duration = duration
	return sv
}

func (sv *SendVideo[ID]) SetWidth(width int) *SendVideo[ID] {
	sv.Width = width
	return sv
}

func (sv *SendVideo[ID]) SetHeight(height int) *SendVideo[ID] {
	sv.Height = height
	return sv
}

func (sv *SendVideo[ID]) SetThumbnail(thumbnail string) *SendVideo[ID] {
	sv.Thumbnail = thumbnail
	return sv
}

func (sv *SendVideo[ID]) SetCaption(caption string) *SendVideo[ID] {
	sv.Caption = caption
	return sv
}

func (sv *SendVideo[ID]) SetParseMode(parseMode string) *SendVideo[ID] {
	sv.ParseMode = parseMode
	return sv
}

func (sv *SendVideo[ID]) SetCaptionEntities(captionEntities []message_entity_dto.MessageEntity) *SendVideo[ID] {
	sv.CaptionEntities = captionEntities
	return sv
}

func (sv *SendVideo[ID]) SetHasSpoiler(hasSpoiler bool) *SendVideo[ID] {
	sv.HasSpoiler = hasSpoiler
	return sv
}

func (sv *SendVideo[ID]) SetSupportsStreaming(supportsStreaming bool) *SendVideo[ID] {
	sv.SupportsStreaming = supportsStreaming
	return sv
}

func (sv *SendVideo[ID]) SetDisableNotification(disableNotification bool) *SendVideo[ID] {
	sv.DisableNotification = disableNotification
	return sv
}

func (sv *SendVideo[ID]) SetProtectContent(protectContent bool) *SendVideo[ID] {
	sv.ProtectContent = protectContent
	return sv
}

func (sv *SendVideo[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendVideo[ID] {
	sv.ReplyToMessageId = replyToMessageId
	return sv
}

func (sv *SendVideo[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendVideo[ID] {
	sv.AllowSendingWithoutReply = allowSendingWithoutReply
	return sv
}

func (sv *SendVideo[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendVideo[ID] {
	sv.ReplyMarkup = replyMarkup
	return sv
}

// https://core.telegram.org/bots/api#sendanimation
type SendAnimation[ID string | int64] struct {
	ChatID                   ID                                 `json:"chat_id"`
	Animation                string                             `json:"animation"`
	MessageThreadId          int                                `json:"message_thread_id,omitempty"`
	Duration                 int                                `json:"duration,omitempty"`
	Width                    int                                `json:"width,omitempty"`
	Height                   int                                `json:"height,omitempty"`
	Thumbnail                string                             `json:"thumbnail,omitempty"`
	Caption                  string                             `json:"caption,omitempty"`
	ParseMode                string                             `json:"parse_mode,omitempty"`
	CaptionEntities          []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                               `json:"has_spoiler,omitempty"`
	DisableNotification      bool                               `json:"disable_notification,omitempty"`
	ProtectContent           bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewSendAnimation[ID string | int64](chatID ID, animation string) *SendAnimation[ID] {
	return &SendAnimation[ID]{
		ChatID:    chatID,
		Animation: animation,
	}
}

func (sa *SendAnimation[ID]) SetMessageThreadId(messageThreadId int) *SendAnimation[ID] {
	sa.MessageThreadId = messageThreadId
	return sa
}

func (sa *SendAnimation[ID]) SetDuration(duration int) *SendAnimation[ID] {
	sa.Duration = duration
	return sa
}

func (sa *SendAnimation[ID]) SetWidth(width int) *SendAnimation[ID] {
	sa.Width = width
	return sa
}

func (sa *SendAnimation[ID]) SetHeight(height int) *SendAnimation[ID] {
	sa.Height = height
	return sa
}

func (sa *SendAnimation[ID]) SetThumbnail(thumbnail string) *SendAnimation[ID] {
	sa.Thumbnail = thumbnail
	return sa
}

func (sa *SendAnimation[ID]) SetCaption(caption string) *SendAnimation[ID] {
	sa.Caption = caption
	return sa
}

func (sa *SendAnimation[ID]) SetParseMode(parseMode string) *SendAnimation[ID] {
	sa.ParseMode = parseMode
	return sa
}

func (sa *SendAnimation[ID]) SetCaptionEntities(captionEntities []message_entity_dto.MessageEntity) *SendAnimation[ID] {
	sa.CaptionEntities = captionEntities
	return sa
}

func (sa *SendAnimation[ID]) SetHasSpoiler(hasSpoiler bool) *SendAnimation[ID] {
	sa.HasSpoiler = hasSpoiler
	return sa
}

func (sa *SendAnimation[ID]) SetDisableNotification(disableNotification bool) *SendAnimation[ID] {
	sa.DisableNotification = disableNotification
	return sa
}

func (sa *SendAnimation[ID]) SetProtectContent(protectContent bool) *SendAnimation[ID] {
	sa.ProtectContent = protectContent
	return sa
}

func (sa *SendAnimation[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendAnimation[ID] {
	sa.ReplyToMessageId = replyToMessageId
	return sa
}

func (sa *SendAnimation[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendAnimation[ID] {
	sa.AllowSendingWithoutReply = allowSendingWithoutReply
	return sa
}

func (sa *SendAnimation[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendAnimation[ID] {
	sa.ReplyMarkup = replyMarkup
	return sa
}

// https://core.telegram.org/bots/api#sendvoice
type SendVoice[ID string | int64] struct {
	ChatID                   ID                                 `json:"chat_id"`
	Voice                    string                             `json:"voice"`
	MessageThreadId          int                                `json:"message_thread_id,omitempty"`
	Caption                  string                             `json:"caption,omitempty"`
	ParseMode                string                             `json:"parse_mode,omitempty"`
	CaptionEntities          []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                                `json:"duration,omitempty"`
	DisableNotification      bool                               `json:"disable_notification,omitempty"`
	ProtectContent           bool                               `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard              `json:"reply_markup,omitempty"`
}

func NewSendVoice[ID string | int64](chatID ID, voice string) *SendVoice[ID] {
	return &SendVoice[ID]{
		ChatID: chatID,
		Voice:  voice,
	}
}

func (sv *SendVoice[ID]) SetMessageThreadId(messageThreadId int) *SendVoice[ID] {
	sv.MessageThreadId = messageThreadId
	return sv
}

func (sv *SendVoice[ID]) SetCaption(caption string) *SendVoice[ID] {
	sv.Caption = caption
	return sv
}

func (sv *SendVoice[ID]) SetParseMode(parseMode string) *SendVoice[ID] {
	sv.ParseMode = parseMode
	return sv
}

func (sv *SendVoice[ID]) SetCaptionEntities(captionEntities []message_entity_dto.MessageEntity) *SendVoice[ID] {
	sv.CaptionEntities = captionEntities
	return sv
}

func (sv *SendVoice[ID]) SetDuration(duration int) *SendVoice[ID] {
	sv.Duration = duration
	return sv
}

func (sv *SendVoice[ID]) SetDisableNotification(disableNotification bool) *SendVoice[ID] {
	sv.DisableNotification = disableNotification
	return sv
}

func (sv *SendVoice[ID]) SetProtectContent(protectContent bool) *SendVoice[ID] {
	sv.ProtectContent = protectContent
	return sv
}

func (sv *SendVoice[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendVoice[ID] {
	sv.ReplyToMessageId = replyToMessageId
	return sv
}

func (sv *SendVoice[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendVoice[ID] {
	sv.AllowSendingWithoutReply = allowSendingWithoutReply
	return sv
}

func (sv *SendVoice[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendVoice[ID] {
	sv.ReplyMarkup = replyMarkup
	return sv
}

// https://core.telegram.org/bots/api#senddice
type SendDice[ID string | int64] struct {
	ChatID                   ID                    `json:"chat_id"`
	Emoji                    string                `json:"voice,omitempty"`
	MessageThreadId          int                   `json:"message_thread_id,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard `json:"reply_markup,omitempty"`
}

func NewSendDice[ID string | int64](chatID ID) *SendDice[ID] {
	return &SendDice[ID]{
		ChatID: chatID,
	}
}

func (s *SendDice[ID]) SetEmoji(emoji string) *SendDice[ID] {
	s.Emoji = emoji
	return s
}

func (s *SendDice[ID]) SetMessageThreadId(messageThreadId int) *SendDice[ID] {
	s.MessageThreadId = messageThreadId
	return s
}

func (s *SendDice[ID]) SetDisableNotification(disableNotification bool) *SendDice[ID] {
	s.DisableNotification = disableNotification
	return s
}

func (s *SendDice[ID]) SetProtectContent(protectContent bool) *SendDice[ID] {
	s.ProtectContent = protectContent
	return s
}

func (s *SendDice[ID]) SetReplyToMessageId(replyToMessageId uint64) *SendDice[ID] {
	s.ReplyToMessageId = replyToMessageId
	return s
}

func (s *SendDice[ID]) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) *SendDice[ID] {
	s.AllowSendingWithoutReply = allowSendingWithoutReply
	return s
}

func (s *SendDice[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendDice[ID] {
	s.ReplyMarkup = replyMarkup
	return s
}

// https://core.telegram.org/bots/api#sendchataction
type SendChatAction[ID string | int64] struct {
	ChatID          ID     `json:"chat_id"`
	MessageThreadId int    `json:"message_thread_id,omitempty"`
	Action          string `json:"caption"`
}

func NewSendChatAction[ID string | int64](chatID ID, action string) *SendChatAction[ID] {
	return &SendChatAction[ID]{
		ChatID: chatID,
		Action: action,
	}
}

func (s *SendChatAction[ID]) SetMessageThreadId(messageThreadId int) *SendChatAction[ID] {
	s.MessageThreadId = messageThreadId
	return s
}

func (s *SendChatAction[ID]) SetAction(action string) *SendChatAction[ID] {
	s.Action = action
	return s
}
