package method_dto

import (
	message_dto "TGoBot/dto/message"
	stickers_dto "TGoBot/dto/stickers"
)

// https://core.telegram.org/bots/api#sendsticker
type SendSticker[ID int64 | string] struct {
	ChatID                   ID                    `json:"chat_id"`
	MessageThreadId          int                   `json:"message_thread_id,omitempty"`
	Sticker                  string                `json:"sticker"`
	Emoji                    string                `json:"emoji,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId         uint64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              message_dto.IKeyboard `json:"reply_markup,omitempty"`
}

func NewSendSticker[ID int64 | string](chatID ID, sticker string) *SendSticker[ID] {
	return &SendSticker[ID]{
		ChatID:  chatID,
		Sticker: sticker,
	}
}

func (ss *SendSticker[ID]) SetMessageThreadId(messageThreadId int) *SendSticker[ID] {
	ss.MessageThreadId = messageThreadId
	return ss
}

func (ss *SendSticker[ID]) SetEmoji(emoji string) *SendSticker[ID] {
	ss.Emoji = emoji
	return ss
}

func (ss *SendSticker[ID]) SetDisableNotification(disable bool) *SendSticker[ID] {
	ss.DisableNotification = disable
	return ss
}

func (ss *SendSticker[ID]) SetProtectContent(protect bool) *SendSticker[ID] {
	ss.ProtectContent = protect
	return ss
}

func (ss *SendSticker[ID]) SetReplyToMessage(replyToMessageID uint64) *SendSticker[ID] {
	ss.ReplyToMessageId = replyToMessageID
	return ss
}

func (ss *SendSticker[ID]) SetAllowSendingWithoutReply(allow bool) *SendSticker[ID] {
	ss.AllowSendingWithoutReply = allow
	return ss
}

func (ss *SendSticker[ID]) SetReplyMarkup(replyMarkup message_dto.IKeyboard) *SendSticker[ID] {
	ss.ReplyMarkup = replyMarkup
	return ss
}

// https://core.telegram.org/bots/api#getstickerset
type GetStickerSet struct {
	Name string `json:"name"`
}

func NewGetStickerSet(name string) *GetStickerSet {
	return &GetStickerSet{
		Name: name,
	}
}

// https://core.telegram.org/bots/api#getcustomemojistickers
type GetCustomEmojiStickers struct {
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

func NewGetCustomEmojiStickers(customEmojiIDs []string) *GetCustomEmojiStickers {
	return &GetCustomEmojiStickers{
		CustomEmojiIDs: customEmojiIDs,
	}
}

// https://core.telegram.org/bots/api#uploadstickerfile
type UploadStickerFile struct {
	UserID        int64  `json:"user_id"`
	StickerFile   string `json:"sticker_file"`
	StickerFormat string `json:"sticker_format,omitempty"`
}

func NewUploadStickerFile(userID int64, stickerFile string) *UploadStickerFile {
	return &UploadStickerFile{
		UserID:      userID,
		StickerFile: stickerFile,
	}
}

func (u *UploadStickerFile) SetStickerFormat(stickerFormat string) *UploadStickerFile {
	u.StickerFormat = stickerFormat
	return u
}

// https://core.telegram.org/bots/api#createnewstickerset
type CreateNewStickerSet struct {
	UserID          int64                       `json:"user_id"`
	Name            string                      `json:"name"`
	Title           string                      `json:"title"`
	Stickers        []stickers_dto.InputSticker `json:"stickers"`
	StickerFormat   string                      `json:"sticker_format"`
	StickerType     string                      `json:"sticker_type,omitempty"`
	NeedsRepainting bool                        `json:"needs_repainting,omitempty"`
}

func NewCreateNewStickerSet(userID int64, name string, title string, stickers []stickers_dto.InputSticker, stickerFormat string) *CreateNewStickerSet {
	return &CreateNewStickerSet{
		UserID:        userID,
		Name:          name,
		Title:         title,
		Stickers:      stickers,
		StickerFormat: stickerFormat,
	}
}

func (c *CreateNewStickerSet) SetStickerType(stickerType string) *CreateNewStickerSet {
	c.StickerType = stickerType
	return c
}

func (c *CreateNewStickerSet) SetNeedsRepainting(needsRepainting bool) *CreateNewStickerSet {
	c.NeedsRepainting = needsRepainting
	return c
}

// https://core.telegram.org/bots/api#addstickertoset
type AddStickerToSet struct {
	UserID  int64                     `json:"user_id"`
	Name    string                    `json:"name"`
	Sticker stickers_dto.InputSticker `json:"sticker"`
}

func NewAddStickerToSet(userID int64, name string, sticker stickers_dto.InputSticker) *AddStickerToSet {
	return &AddStickerToSet{
		UserID:  userID,
		Name:    name,
		Sticker: sticker,
	}
}

// https://core.telegram.org/bots/api#setstickerpositioninset
type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

func NewSetStickerPositionInSet(sticker string, position int) *SetStickerPositionInSet {
	return &SetStickerPositionInSet{
		Sticker:  sticker,
		Position: position,
	}
}

// https://core.telegram.org/bots/api#deletestickerfromset
type DeleteStickerFromSet struct {
	Sticker string `json:"sticker"`
}

func NewDeleteStickerFromSet(sticker string) *DeleteStickerFromSet {
	return &DeleteStickerFromSet{
		Sticker: sticker,
	}
}

// https://core.telegram.org/bots/api#setstickeremojilist
type SetStickerEmojiList struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

func NewSetStickerEmojiList(sticker string, emojiList []string) *SetStickerEmojiList {
	return &SetStickerEmojiList{
		Sticker:   sticker,
		EmojiList: emojiList,
	}
}

// https://core.telegram.org/bots/api#setstickerkeywords
type SetStickerKeywords struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords,omitempty"`
}

func NewSetStickerKeywords(sticker string) *SetStickerKeywords {
	return &SetStickerKeywords{
		Sticker: sticker,
	}
}

func (ssk *SetStickerKeywords) SetKeywords(keywords []string) *SetStickerKeywords {
	ssk.Keywords = keywords
	return ssk
}

// https://core.telegram.org/bots/api#setstickermaskposition
type SetStickerMaskPosition struct {
	Sticker      string                     `json:"sticker"`
	MaskPosition *stickers_dto.MaskPosition `json:"mask_position,omitempty"`
}

func NewSetStickerMaskPosition(sticker string) *SetStickerMaskPosition {
	return &SetStickerMaskPosition{
		Sticker: sticker,
	}
}

func (ssmp *SetStickerMaskPosition) SetMaskPosition(maskPosition *stickers_dto.MaskPosition) *SetStickerMaskPosition {
	ssmp.MaskPosition = maskPosition
	return ssmp
}

// https://core.telegram.org/bots/api#setstickersettitle
type SetStickerSetTitle struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func NewSetStickerSetTitle(name, title string) *SetStickerSetTitle {
	return &SetStickerSetTitle{
		Name:  name,
		Title: title,
	}
}

// https://core.telegram.org/bots/api#setstickersetthumbnail
type SetStickerSetThumbnail struct {
	Name      string `json:"name"`
	UserID    int    `json:"user_id"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

func NewSetStickerSetThumbnail(name string, userID int) *SetStickerSetThumbnail {
	return &SetStickerSetThumbnail{
		Name:   name,
		UserID: userID,
	}
}

func (sst *SetStickerSetThumbnail) SetThumbnail(thumbnail string) *SetStickerSetThumbnail {
	sst.Thumbnail = thumbnail
	return sst
}

// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
type SetCustomEmojiStickerSetThumbnail struct {
	Name          string `json:"name"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

func NewSetCustomEmojiStickerSetThumbnail(name string) *SetCustomEmojiStickerSetThumbnail {
	return &SetCustomEmojiStickerSetThumbnail{
		Name: name,
	}
}

func (scest *SetCustomEmojiStickerSetThumbnail) SetCustomEmojiID(customEmojiID string) *SetCustomEmojiStickerSetThumbnail {
	scest.CustomEmojiID = customEmojiID
	return scest
}

// https://core.telegram.org/bots/api#deletestickerset
type DeleteStickerSet struct {
	Name string `json:"name"`
}

func NewDeleteStickerSet(name string) *DeleteStickerSet {
	return &DeleteStickerSet{
		Name: name,
	}
}
