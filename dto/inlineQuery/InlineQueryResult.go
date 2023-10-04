package inline_dto

import message_dto "github.com/Avirtan/TGoBot/dto/message"

// TODO Сделать билдер для каждого класса, и сделать создание резеультатов черезе New, NewТИП_РЕЗУЛЬТАТА параметры в New будут все, кроме omitempty
// для omitempty параметров делать методы и возвращать сам объект, тем самы получится билдер

// https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult interface {
	GetType() string
}

// https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	Type               string              `json:"type"`
	Id                 string              `json:"id"`
	Title              string              `json:"title"`
	InputMssageContent InputMessageContent `json:"input_message_content"`
	Url                string              `json:"url,omitempty"`
	HideUrl            bool                `json:"hide_url,omitempty"`
	Description        string              `json:"description,omitempty"`
	ThumbnailUrl       string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth     int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight    int                 `json:"thumbnail_height,omitempty"`
}

func (q *InlineQueryResultArticle) GetType() string {
	return "article"
}

// https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	PhotoUrl            string                            `json:"photo_url"`
	ThumbnailUrl        string                            `json:"thumbnail_url"`
	PhotoWidth          int                               `json:"photo_width,omitempty"`
	PhotoHeight         int                               `json:"photo_height,omitempty"`
	Title               string                            `json:"title,omitempty"`
	Description         string                            `json:"description,omitempty"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []*message_dto.MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultPhoto) GetType() string {
	return "photo"
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	PhotoFileId         string                            `json:"photo_file_id"`
	Title               string                            `json:"title,omitempty"`
	Description         string                            `json:"description,omitempty"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []*message_dto.MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedPhoto) GetType() string {
	return "photo"
}

// https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	GifUrl              string                            `json:"gif_url"`
	GifWidth            int                               `json:"gif_width,omitempty"`
	GifHeight           int                               `json:"gif_height,omitempty"`
	GifDuration         int                               `json:"gif_duration,omitempty"`
	ThumbnailUrl        string                            `json:"thumbnail_url,omitempty"`
	ThumbnailMimeType   string                            `json:"thumbnail_mime_type,omitempty"`
	Title               string                            `json:"title,omitempty"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []*message_dto.MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultGif) GetType() string {
	return "gif"
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	GifFileId           string                            `json:"gif_file_id"`
	Title               string                            `json:"title,omitempty"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []*message_dto.MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedGif) GetType() string {
	return "gif"
}

// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMpeg4Gif struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	Mpeg4Url            string                            `json:"mpeg4_url"`
	Mpeg4Width          int                               `json:"gif_width,omitempty"`
	Mpeg4Height         int                               `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                               `json:"mpeg4_duration,omitempty"`
	ThumbnailUrl        string                            `json:"thumbnail_url,omitempty"`
	ThumbnailMimeType   string                            `json:"thumbnail_mime_type,omitempty"`
	Title               string                            `json:"title,omitempty"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []*message_dto.MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultMpeg4Gif) GetType() string {
	return "mpeg4_gif"
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	Mpeg4FileId         string                            `json:"mpeg4_file_id"`
	Title               string                            `json:"title,omitempty"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []*message_dto.MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedMpeg4Gif) GetType() string {
	return "mpeg4_gif"
}

// https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	VideoUrl            string                            `json:"video_url"`
	MimeType            string                            `json:"mime_type"`
	ThumbnailUrl        string                            `json:"thumbnail_url"`
	Title               string                            `json:"title"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	VideoWidth          int                               `json:"video_width,omitempty"`
	VideoHeight         int                               `json:"video_height,omitempty"`
	VideoDuration       int                               `json:"video_duration"`
	Description         string                            `json:"description,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultVideo) GetType() string {
	return "video"
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	VideoFileId         string                            `json:"video_file_id"`
	Title               string                            `json:"title"`
	Description         string                            `json:"description,omitempty"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedVideo) GetType() string {
	return "video"
}

// https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	AudioUrl            string                            `json:"audio_url"`
	Title               string                            `json:"title"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	Performer           string                            `json:"performer,omitempty"`
	AudioDuration       int                               `json:"audio_duration,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultAudio) GetType() string {
	return "audio"
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	AudioFileId         string                            `json:"audio_file_id"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedAudio) GetType() string {
	return "audio"
}

// https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	VoiceUrl            string                            `json:"voice_url"`
	Title               string                            `json:"title"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultVoice) GetType() string {
	return "voice"
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	VoiceFileId         string                            `json:"voice_file_id"`
	Title               string                            `json:"title"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedVoice) GetType() string {
	return "voice"
}

// https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	Title               string                            `json:"title"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	DocumentUrl         string                            `json:"document_url"`
	MimeType            string                            `json:"mime_type"`
	Description         string                            `json:"description,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
	ThumbnailUrl        string                            `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                               `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                               `json:"thumbnail_height,omitempty"`
}

func (q *InlineQueryResultDocument) GetType() string {
	return "document"
}

// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	Title               string                            `json:"title"`
	DocumentFileId      string                            `json:"document_file_id"`
	Caption             string                            `json:"caption,omitempty"`
	ParseMode           string                            `json:"parse_mode,omitempty"`
	CaptionEntities     []message_dto.MessageEntity       `json:"caption_entities,omitempty"`
	Description         string                            `json:"description,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedDocument) GetType() string {
	return "document"
}

// https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	Type                 string                            `json:"type"`
	Id                   string                            `json:"id"`
	Latitude             float64                           `json:"latitude"`
	Longitude            float64                           `json:"longitude"`
	Title                string                            `json:"title"`
	HorizontalAccuracy   float64                           `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int32                             `json:"live_period,omitempty"`
	Heading              uint16                            `json:"heading,omitempty"`
	ProximityAlertRadius uint32                            `json:"Proximity_alert_radius,omitempty"`
	ReplyMarkup          *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  InputMessageContent               `json:"input_message_content,omitempty"`
	ThumbnailUrl         string                            `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       uint                              `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      uint                              `json:"thumbnail_height,omitempty"`
}

func (q *InlineQueryResultLocation) GetType() string {
	return "location"
}

// https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	Latitude            float64                           `json:"latitude"`
	Longitude           float64                           `json:"longitude"`
	Title               string                            `json:"title"`
	Address             string                            `json:"address"`
	FoursquareId        string                            `json:"foursquare_id,omitempty"`
	FoursquareType      string                            `json:"foursquare_type,omitempty"`
	GooglePlaceId       string                            `json:"google_place_id,omitempty"`
	GooglePlaceType     string                            `json:"google_place_type,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
	ThumbnailUrl        string                            `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      uint                              `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     uint                              `json:"thumbnail_height,omitempty"`
}

func (q *InlineQueryResultVenue) GetType() string {
	return "venue"
}

// https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	PhoneNumber         string                            `json:"phone_number"`
	FirstName           string                            `json:"first_name"`
	LastName            string                            `json:"last_name,omitempty"`
	Vcard               string                            `json:"vcard,omitempty"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
	ThumbnailUrl        string                            `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      uint                              `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     uint                              `json:"thumbnail_height,omitempty"`
}

func (q *InlineQueryResultContact) GetType() string {
	return "contact"
}

// https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultGame struct {
	Type          string                            `json:"type"`
	Id            string                            `json:"id"`
	GameShortName string                            `json:"game_short_name"`
	ReplyMarkup   *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (q *InlineQueryResultGame) GetType() string {
	return "game"
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	Type                string                            `json:"type"`
	Id                  string                            `json:"id"`
	StickerFileId       string                            `json:"sticker_file_id"`
	ReplyMarkup         *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent               `json:"input_message_content,omitempty"`
}

func (q *InlineQueryResultCachedSticker) GetType() string {
	return "sticker"
}
