package method_dto

import (
	message_dto "github.com/Avirtan/goborg-tg/dto/message"
	payment_dto "github.com/Avirtan/goborg-tg/dto/payment"
)

// https://core.telegram.org/bots/api#sendinvoice
type SendInvoice struct {
	ChatID                    int64                             `json:"chat_id"`
	Title                     string                            `json:"title"`
	Description               string                            `json:"description"`
	Payload                   string                            `json:"payload"`
	ProviderToken             string                            `json:"provider_token"`
	Currency                  string                            `json:"currency"`
	Prices                    []payment_dto.LabeledPrice        `json:"prices"`
	MaxTipAmount              int                               `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                             `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                            `json:"start_parameter,omitempty"`
	ProviderData              string                            `json:"provider_data,omitempty"`
	PhotoURL                  string                            `json:"photo_url,omitempty"`
	PhotoSize                 int                               `json:"photo_size,omitempty"`
	PhotoWidth                int                               `json:"photo_width,omitempty"`
	PhotoHeight               int                               `json:"photo_height,omitempty"`
	NeedName                  bool                              `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                              `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                              `json:"need_email,omitempty"`
	NeedShippingAddress       bool                              `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                              `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                              `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                              `json:"is_flexible,omitempty"`
	DisableNotification       bool                              `json:"disable_notification,omitempty"`
	ProtectContent            bool                              `json:"protect_content,omitempty"`
	ReplyToMessageID          int                               `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply  bool                              `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup               *message_dto.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func NewSendInvoice(chatID int64, title, description, payload, providerToken, currency string, prices []payment_dto.LabeledPrice) *SendInvoice {
	return &SendInvoice{
		ChatID:        chatID,
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}
}

func (s *SendInvoice) SetMaxTipAmount(maxTipAmount int) {
	s.MaxTipAmount = maxTipAmount
}

func (s *SendInvoice) SetSuggestedTipAmounts(suggestedTipAmounts []int) {
	s.SuggestedTipAmounts = suggestedTipAmounts
}

func (s *SendInvoice) SetStartParameter(startParameter string) {
	s.StartParameter = startParameter
}

func (s *SendInvoice) SetProviderData(providerData string) {
	s.ProviderData = providerData
}

func (s *SendInvoice) SetPhoto(photoURL string, photoSize int, photoWidth int, photoHeight int) {
	s.PhotoURL = photoURL
	s.PhotoSize = photoSize
	s.PhotoWidth = photoWidth
	s.PhotoHeight = photoHeight
}

func (s *SendInvoice) SetNeedName(needName bool) {
	s.NeedName = needName
}

func (s *SendInvoice) SetNeedPhoneNumber(needPhoneNumber bool) {
	s.NeedPhoneNumber = needPhoneNumber
}

func (s *SendInvoice) SetNeedEmail(needEmail bool) {
	s.NeedEmail = needEmail
}

func (s *SendInvoice) SetNeedShippingAddress(needShippingAddress bool) {
	s.NeedShippingAddress = needShippingAddress
}

func (s *SendInvoice) SetSendPhoneNumberToProvider(sendPhoneNumberToProvider bool) {
	s.SendPhoneNumberToProvider = sendPhoneNumberToProvider
}

func (s *SendInvoice) SetSendEmailToProvider(sendEmailToProvider bool) {
	s.SendEmailToProvider = sendEmailToProvider
}

func (s *SendInvoice) SetIsFlexible(isFlexible bool) {
	s.IsFlexible = isFlexible
}

func (s *SendInvoice) SetDisableNotification(disableNotification bool) {
	s.DisableNotification = disableNotification
}

func (s *SendInvoice) SetProtectContent(protectContent bool) {
	s.ProtectContent = protectContent
}

func (s *SendInvoice) SetReplyToMessageID(replyToMessageID int) {
	s.ReplyToMessageID = replyToMessageID
}

func (s *SendInvoice) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) {
	s.AllowSendingWithoutReply = allowSendingWithoutReply
}

func (s *SendInvoice) SetReplyMarkup(replyMarkup *message_dto.InlineKeyboardMarkup) {
	s.ReplyMarkup = replyMarkup
}

// https://core.telegram.org/bots/api#createinvoicelink
type CreateInvoiceLink struct {
	Title               string                     `json:"title"`
	Description         string                     `json:"description"`
	Payload             string                     `json:"payload"`
	ProviderToken       string                     `json:"provider_token"`
	Currency            string                     `json:"currency"`
	Prices              []payment_dto.LabeledPrice `json:"prices"`
	MaxTipAmount        int                        `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts []int                      `json:"suggested_tip_amounts,omitempty"`
	ProviderData        string                     `json:"provider_data,omitempty"`
	PhotoURL            string                     `json:"photo_url,omitempty"`
	PhotoSize           int                        `json:"photo_size,omitempty"`
	PhotoWidth          int                        `json:"photo_width,omitempty"`
	PhotoHeight         int                        `json:"photo_height,omitempty"`
	NeedName            bool                       `json:"need_name,omitempty"`
	NeedPhoneNumber     bool                       `json:"need_phone_number,omitempty"`
	NeedEmail           bool                       `json:"need_email,omitempty"`
	NeedShippingAddr    bool                       `json:"need_shipping_address,omitempty"`
	SendPhoneToProvider bool                       `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider bool                       `json:"send_email_to_provider,omitempty"`
	IsFlexible          bool                       `json:"is_flexible,omitempty"`
}

func NewCreateInvoiceLink(title, description, payload, providerToken, currency string, prices []payment_dto.LabeledPrice) *CreateInvoiceLink {
	return &CreateInvoiceLink{
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}
}

func (cil *CreateInvoiceLink) SetMaxTipAmount(maxTipAmount int) {
	cil.MaxTipAmount = maxTipAmount
}

func (cil *CreateInvoiceLink) SetSuggestedTipAmounts(suggestedTipAmounts []int) {
	cil.SuggestedTipAmounts = suggestedTipAmounts
}

func (cil *CreateInvoiceLink) SetProviderData(providerData string) {
	cil.ProviderData = providerData
}

func (cil *CreateInvoiceLink) SetPhotoURL(photoURL string) {
	cil.PhotoURL = photoURL
}

func (cil *CreateInvoiceLink) SetPhotoSize(photoSize int) {
	cil.PhotoSize = photoSize
}

func (cil *CreateInvoiceLink) SetPhotoDimensions(width, height int) {
	cil.PhotoWidth = width
	cil.PhotoHeight = height
}

func (cil *CreateInvoiceLink) SetNeedName(needName bool) {
	cil.NeedName = needName
}

func (cil *CreateInvoiceLink) SetNeedPhoneNumber(needPhoneNumber bool) {
	cil.NeedPhoneNumber = needPhoneNumber
}

func (cil *CreateInvoiceLink) SetNeedEmail(needEmail bool) {
	cil.NeedEmail = needEmail
}

func (cil *CreateInvoiceLink) SetNeedShippingAddress(needShippingAddress bool) {
	cil.NeedShippingAddr = needShippingAddress
}

func (cil *CreateInvoiceLink) SetSendPhoneToProvider(sendPhoneToProvider bool) {
	cil.SendPhoneToProvider = sendPhoneToProvider
}

func (cil *CreateInvoiceLink) SetSendEmailToProvider(sendEmailToProvider bool) {
	cil.SendEmailToProvider = sendEmailToProvider
}

func (cil *CreateInvoiceLink) SetIsFlexible(isFlexible bool) {
	cil.IsFlexible = isFlexible
}

// https://core.telegram.org/bots/api#answershippingquery
type AnswerShippingQuery struct {
	ShippingQueryID string                       `json:"shipping_query_id"`
	OK              bool                         `json:"ok"`
	ShippingOptions []payment_dto.ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string                       `json:"error_message,omitempty"`
}

func NewAnswerShippingQuery(shippingQueryID string, ok bool) *AnswerShippingQuery {
	return &AnswerShippingQuery{
		ShippingQueryID: shippingQueryID,
		OK:              ok,
	}
}

func (asq *AnswerShippingQuery) SetShippingOptions(shippingOptions []payment_dto.ShippingOption) {
	asq.ShippingOptions = shippingOptions
}

func (asq *AnswerShippingQuery) SetErrorMessage(errorMessage string) {
	asq.ErrorMessage = errorMessage
}

// https://core.telegram.org/bots/api#answerprecheckoutquery
type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	OK                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

func NewAnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool) *AnswerPreCheckoutQuery {
	return &AnswerPreCheckoutQuery{
		PreCheckoutQueryID: preCheckoutQueryID,
		OK:                 ok,
	}
}

func (apcq *AnswerPreCheckoutQuery) SetErrorMessage(errorMessage string) {
	apcq.ErrorMessage = errorMessage
}
