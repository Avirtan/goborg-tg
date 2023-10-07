package inline_dto

import message_entity_dto "github.com/Avirtan/TGoBot/dto/message_entity"

// https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent interface {
	InputMessage()
}

// https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText           string                             `json:"message_text"`
	ParseMode             string                             `json:"parse_mode,omitempty"`
	Entities              []message_entity_dto.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                               `json:"disable_web_page_preview,omitempty"`
}

func (m *InputTextMessageContent) InputMessage() {

}

// https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           uint32  `json:"live_period,omitempty"`
	Heading              uint16  `json:"heading,omitempty"`
	ProximityAlertRadius uint32  `json:"proximity_alert_radius,omitempty"`
}

// https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"horizontal_actitlecuracy"`
	Address         string  `json:"address"`
	FoursquareId    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceId   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

// https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"foursquare_id,omitempty"`
	Vcard       string `json:"foursquare_type,omitempty"`
}

// https://core.telegram.org/bots/api#inputinvoicemessagecontent
type InputInvoiceMessageContent struct {
	Title                     string `json:"title"`
	Description               string `json:"description"`
	Payload                   string `json:"payload"`
	ProviderToken             string `json:"provider_token"`
	Currency                  string `json:"currency"`
	Prices                    []any  `json:"prices"` // TODO LabeledPrice
	MaxTipAmount              uint   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []uint `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string `json:"provider_data,omitempty"`
	PhotoUrl                  string `json:"photo_url,omitempty"`
	PhotoSize                 uint   `json:"photo_size,omitempty"`
	PhotoWidth                uint   `json:"photo_width,omitempty"`
	PhotoHeight               uint   `json:"photo_height,omitempty"`
	NeedName                  bool   `json:"need_name,omitempty"`
	NeedPhoneNumber           bool   `json:"need_phone_number,omitempty"`
	NeedEmail                 bool   `json:"need_email,omitempty"`
	NeedShippingAddress       bool   `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool   `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool   `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool   `json:"is_flexible,omitempty"`
}
