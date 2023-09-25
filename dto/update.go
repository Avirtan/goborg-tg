package dto

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Update []Update `json:"result"`
}

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID           uint64            `json:"update_id"`
	Message            *Message          `json:"message,omitempty"`
	EditedMessage      *Message          `json:"edited_message,omitempty"`
	ChannelPost        *Message          `json:"channel_post,omitempty"`
	EditedChannelPost  *Message          `json:"edited_channel_post,omitempty"`
	InlineQuery        any               `json:"inline_query,omitempty"`         // TODO InlineQuery
	ChosenInlineResult any               `json:"chosen_inline_result,omitempty"` // TODO ChosenInlineResult
	CallbackQuery      CallbackQuery     `json:"callback_query,omitempty"`
	ShippingQuery      any               `json:"shipping_query,omitempty"`     // TODO ShippingQuery
	PreCheckoutQuery   any               `json:"pre_checkout_query,omitempty"` // TODO PreCheckoutQuery
	Poll               Poll              `json:"poll,omitempty"`
	PollAnswer         PollAnswer        `json:"poll_answer,omitempty"`
	MyChatMember       ChatMemberUpdated `json:"my_chat_member,omitempty"`
	ChatMember         ChatMemberUpdated `json:"chat_member,omitempty"`
	ChatJoinRequest    ChatJoinRequest   `json:"chat_join_request,omitempty"`
}

// https://core.telegram.org/bots/api#getupdates
type UpdateRequest struct {
	Offset  uint64 `json:"offset"`
	Limit   int    `json:"limmit"`
	Timeout int    `json:"timeout"`
}
