package message_dto

import (
	forum_dto "github.com/Avirtan/goborg-tg/dto/forum"
	game_dto "github.com/Avirtan/goborg-tg/dto/game"
	message_entity_dto "github.com/Avirtan/goborg-tg/dto/message_entity"
	passport_dto "github.com/Avirtan/goborg-tg/dto/passport"
	payment_dto "github.com/Avirtan/goborg-tg/dto/payment"
	stickers_dto "github.com/Avirtan/goborg-tg/dto/stickers"
	user_dto "github.com/Avirtan/goborg-tg/dto/user"
	utils_dto "github.com/Avirtan/goborg-tg/dto/utils"
	video_dto "github.com/Avirtan/goborg-tg/dto/video"
)

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId                     int                                     `json:"message_id"`
	MessageThreadId               int                                     `json:"message_thread_id"`
	Text                          string                                  `json:"text"`
	From                          user_dto.User                           `json:"from"`
	SenderChat                    Chat                                    `json:"sender_chat"`
	Date                          int                                     `json:"date"`
	Chat                          Chat                                    `json:"chat"`
	ForwardFrom                   *user_dto.User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                                   `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId          int                                     `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                                  `json:"forward_signature,omitempty"`
	ForwardSenderName             string                                  `json:"forward_sender_name,omitempty"`
	ForwardDate                   int                                     `json:"forward_date,omitempty"`
	IsTopicMessage                bool                                    `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                                    `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                                `json:"reply_to_message,omitempty"`
	ViaBot                        *user_dto.User                          `json:"via_bot,omitempty"`
	EditDate                      int                                     `json:"edit_date,omitempty"`
	HasProtectedContent           bool                                    `json:"has_protected_content,omitempty"`
	MediaGroupId                  string                                  `json:"media_group_id,omitempty"`
	AuthorSignature               string                                  `json:"author_signature,omitempty"`
	Entities                      []message_entity_dto.MessageEntity      `json:"entities,omitempty"`
	Animation                     *utils_dto.Animation                    `json:"animation,omitempty"`
	Audio                         *utils_dto.Audio                        `json:"audio,omitempty"`
	Document                      *utils_dto.Document                     `json:"document,omitempty"`
	Photo                         []*utils_dto.PhotoSize                  `json:"photo,omitempty"`
	Sticker                       *stickers_dto.Sticker                   `json:"sticker,omitempty"`
	Story                         any                                     `json:"story,omitempty"`
	Video                         *video_dto.Video                        `json:"video,omitempty"`
	VideoNote                     *video_dto.VideoNote                    `json:"video_note,omitempty"`
	Voice                         *utils_dto.Voice                        `json:"voice,omitempty"`
	Caption                       string                                  `json:"caption,omitempty"`
	CaptionEntities               []message_entity_dto.MessageEntity      `json:"caption_entities,omitempty"`
	HasMediaSpoiler               bool                                    `json:"has_media_spoiler,omitempty"`
	Contact                       *user_dto.Contact                       `json:"contact,omitempty"`
	Dice                          Dice                                    `json:"dice,omitempty"`
	Game                          *game_dto.Game                          `json:"game,omitempty"`
	Poll                          *Poll                                   `json:"poll,omitempty"`
	Venue                         *utils_dto.Venue                        `json:"venue,omitempty"`
	Location                      *utils_dto.Location                     `json:"location,omitempty"`
	NewChatMembers                []*user_dto.User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *user_dto.User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                                  `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*utils_dto.PhotoSize                  `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                                    `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                                    `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                                    `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                                    `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged           `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int                                     `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int                                     `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                                `json:"pinned_message,omitempty"`
	Invoice                       *payment_dto.Invoice                    `json:"invoice,omitempty"`
	SuccessfulPayment             *payment_dto.SuccessfulPayment          `json:"successful_payment,omitempty"`
	UserShared                    *UserShared                             `json:"user_shared,omitempty"`
	ChatShared                    *ChatShared                             `json:"chat_shared,omitempty"`
	ConnectedWebsite              string                                  `json:"connected_website,omitempty"`
	WriteAccessAllowed            *utils_dto.WriteAccessAllowed           `json:"write_access_allowed,omitempty"`
	PassportData                  *passport_dto.PassportData              `json:"passport_data,omitempty"`
	ProximityAlertTriggered       ProximityAlertTriggered                 `json:"proximity_alert_triggered,omitempty"`
	ForumTopicCreated             *forum_dto.ForumTopicCreated            `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *forum_dto.ForumTopicEdited             `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              any                                     `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            any                                     `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       any                                     `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     any                                     `json:"general_forum_topic_unhidden,omitempty"`
	VideoChatScheduled            *video_dto.VideoChatScheduled           `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              any                                     `json:"video_chat_started,omitempty"`
	VideoChatParticipantsInvited  *video_dto.VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *utils_dto.WebAppData                   `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup                   `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}
