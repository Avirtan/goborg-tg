package dto

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId                     int                           `json:"message_id"`
	MessageThreadId               int                           `json:"message_thread_id"`
	Text                          string                        `json:"text"`
	From                          User                          `json:"from"`
	SenderChat                    Chat                          `json:"sender_chat"`
	Date                          int                           `json:"date"`
	Chat                          Chat                          `json:"chat"`
	ForwardFrom                   User                          `json:"forward_from,omitempty"`
	ForwardFromChat               Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId          int                           `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                        `json:"forward_signature,omitempty"`
	ForwardSenderName             string                        `json:"forward_sender_name,omitempty"`
	ForwardDate                   int                           `json:"forward_date,omitempty"`
	IsTopicMessage                bool                          `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                          `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                      `json:"reply_to_message,omitempty"`
	ViaBot                        User                          `json:"via_bot,omitempty"`
	EditDate                      int                           `json:"edit_date,omitempty"`
	HasProtectedContent           bool                          `json:"has_protected_content,omitempty"`
	MediaGroupId                  string                        `json:"media_group_id,omitempty"`
	AuthorSignature               string                        `json:"author_signature,omitempty"`
	Entities                      []MessageEntity               `json:"entities,omitempty"`
	Animation                     Animation                     `json:"animation,omitempty"`
	Audio                         Audio                         `json:"audio,omitempty"`
	Document                      Document                      `json:"document,omitempty"`
	Photo                         []PhotoSize                   `json:"photo,omitempty"`
	Sticker                       any                           `json:"sticker,omitempty"` // TODO Sticker
	Story                         any                           `json:"story,omitempty"`
	Video                         Video                         `json:"video,omitempty"`
	VideoNote                     VideoNote                     `json:"video_note,omitempty"`
	Voice                         Voice                         `json:"voice,omitempty"`
	Caption                       string                        `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity               `json:"caption_entities,omitempty"`
	HasMediaSpoiler               bool                          `json:"has_media_spoiler,omitempty"`
	Contact                       Contact                       `json:"contact,omitempty"`
	Dice                          Dice                          `json:"dice,omitempty"`
	Game                          any                           `json:"game,omitempty"`
	Poll                          Poll                          `json:"poll,omitempty"`
	Venue                         Venue                         `json:"venue,omitempty"`
	Location                      Location                      `json:"location,omitempty"`
	NewChatMembers                []User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                          `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int                           `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int                           `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                      `json:"pinned_message,omitempty"`
	Invoice                       any                           `json:"invoice,omitempty"`            // TODO Invoice
	SuccessfulPayment             any                           `json:"successful_payment,omitempty"` // TODO SuccessfulPayment
	UserShared                    UserShared                    `json:"user_shared,omitempty"`
	ChatShared                    ChatShared                    `json:"chat_shared,omitempty"`
	ConnectedWebsite              string                        `json:"connected_website,omitempty"`
	WriteAccessAllowed            WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  any                           `json:"passport_data,omitempty"` // TODO PassportData
	ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	ForumTopicCreated             ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              any                           `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            any                           `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       any                           `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     any                           `json:"general_forum_topic_unhidden,omitempty"`
	VideoChatScheduled            VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              any                           `json:"video_chat_started,omitempty"`
	VideoChatParticipantsInvited  VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
	User          User   `json:"user,omitempty"`
}

// https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type SendMessage struct {
	ChatID      int64    `json:"chat_id"`
	Text        string   `json:"text"`
	ReplyMarkup Keyboard `json:"reply_markup,omitempty"`
}
