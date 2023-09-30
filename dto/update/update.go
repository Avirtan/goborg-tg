package update_dto

import (
	inline_dto "TGoBot/dto/inlineQuery"
	message_dto "TGoBot/dto/message"
)

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Update []Update `json:"result"`
}

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID           uint64                         `json:"update_id"`
	Message            *message_dto.Message           `json:"message,omitempty"`
	EditedMessage      *message_dto.Message           `json:"edited_message,omitempty"`
	ChannelPost        *message_dto.Message           `json:"channel_post,omitempty"`
	EditedChannelPost  *message_dto.Message           `json:"edited_channel_post,omitempty"`
	InlineQuery        *inline_dto.InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *inline_dto.ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *message_dto.CallbackQuery     `json:"callback_query,omitempty"`
	ShippingQuery      *any                           `json:"shipping_query,omitempty"`     // TODO ShippingQuery
	PreCheckoutQuery   *any                           `json:"pre_checkout_query,omitempty"` // TODO PreCheckoutQuery
	Poll               *message_dto.Poll              `json:"poll,omitempty"`
	PollAnswer         *message_dto.PollAnswer        `json:"poll_answer,omitempty"`
	MyChatMember       *message_dto.ChatMemberUpdated `json:"my_chat_member,omitempty"`
	ChatMember         *message_dto.ChatMemberUpdated `json:"chat_member,omitempty"`
	ChatJoinRequest    *message_dto.ChatJoinRequest   `json:"chat_join_request,omitempty"`
}

// https://core.telegram.org/bots/api#getupdates
type UpdateRequest struct {
	Offset         uint64    `json:"offset,omitempty"`
	Limit          int       `json:"limmit,omitempty"`
	Timeout        int       `json:"timeout,omitempty"`
	AllowedUpdates []*string `json:"allowed_updates,omitempty"`
}
