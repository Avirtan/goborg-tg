package method_dto

import message_dto "TGoBot/dto/message"

// https://core.telegram.org/bots/api#sendphoto
type SendPhoto[ID string | int64] struct {
	ChatID                   ID                          `json:"chat_id"`
	Photo                    string                      `json:"photo"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	Caption                  string                      `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	CaptionEntities          []message_dto.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                        `json:"has_spoiler,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendaudio
type SendAudio[ID string | int64] struct {
	ChatID                   ID                          `json:"chat_id"`
	Audio                    string                      `json:"audio"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	Caption                  string                      `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	CaptionEntities          []message_dto.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                         `json:"duration,omitempty"`
	Performer                string                      `json:"performer,omitempty"`
	Title                    string                      `json:"title,omitempty"`
	Thumbnail                string                      `json:"thumbnail,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#senddocument
type SendDocument[ID string | int64] struct {
	ChatID                      ID                          `json:"chat_id"`
	Document                    string                      `json:"document"`
	Thumbnail                   string                      `json:"thumbnail,omitempty"`
	MessageThreadId             int                         `json:"message_thread_id,omitempty"`
	Caption                     string                      `json:"caption,omitempty"`
	ParseMode                   string                      `json:"parse_mode,omitempty"`
	CaptionEntities             []message_dto.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                        `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                        `json:"disable_notification,omitempty"`
	ProtectContent              bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId            uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvideo
type SendVideo[ID string | int64] struct {
	ChatID                   ID                          `json:"chat_id"`
	Video                    string                      `json:"video"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	Duration                 int                         `json:"duration,omitempty"`
	Width                    int                         `json:"width,omitempty"`
	Height                   int                         `json:"height,omitempty"`
	Thumbnail                string                      `json:"thumbnail,omitempty"`
	Caption                  string                      `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	CaptionEntities          []message_dto.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                        `json:"has_spoiler,omitempty"`
	SupportsStreaming        bool                        `json:"supports_streaming,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendanimation
type SendAnimation[ID string | int64] struct {
	ChatID                   ID                          `json:"chat_id"`
	Animation                string                      `json:"animation"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	Duration                 int                         `json:"duration,omitempty"`
	Width                    int                         `json:"width,omitempty"`
	Height                   int                         `json:"height,omitempty"`
	Thumbnail                string                      `json:"thumbnail,omitempty"`
	Caption                  string                      `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	CaptionEntities          []message_dto.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                        `json:"has_spoiler,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                      `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard       `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#sendvoice
type SendVoice[ID string | int64] struct {
	ChatID                   ID                          `json:"chat_id"`
	Voice                    string                      `json:"voice"`
	MessageThreadId          int                         `json:"message_thread_id,omitempty"`
	Caption                  string                      `json:"caption,omitempty"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	CaptionEntities          []message_dto.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                         `json:"duration,omitempty"`
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
