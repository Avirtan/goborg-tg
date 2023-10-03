package payment_dto

import user_dto "TGoBot/dto/user"

// https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state,omitempty"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2,omitempty"`
	PostCode    string `json:"post_code"`
}

// https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *user_dto.User   `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	ID               string         `json:"id"`
	From             *user_dto.User `json:"from"`
	Currency         string         `json:"currency"`
	TotalAmount      int            `json:"total_amount"`
	InvoicePayload   string         `json:"invoice_payload"`
	ShippingOptionID string         `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo     `json:"order_info,omitempty"`
}
