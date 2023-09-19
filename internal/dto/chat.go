package dto

// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id                                 int64        `json:"id"`
	Type                               string       `json:"type"`
	Title                              string       `json:"title"`
	Username                           string       `json:"username"`
	FirstName                          string       `json:"first_name"`
	LastName                           string       `json:"last_name"`
	IsForum                            bool         `json:"is_forum"`
	Photo                              ChatPhoto    `json:"photo"`
	ActiveUsernames                    []string     `json:"active_usernames"`
	EmojiStatusCustomEmojiId           string       `json:"added_to_attachment_menu"`
	EmojiStatusExpirationDate          int          `json:"emoji_status_expiration_date"`
	Bio                                string       `json:"bio"`
	HasPrivateForwards                 bool         `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages bool         `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 bool         `json:"join_to_send_messages"`
	JoinByRequest                      bool         `json:"join_by_request"`
	Description                        string       `json:"description"`
	InviteLink                         string       `json:"invite_link"`
	PinnedMessage                      *Message     `json:"pinned_message"`
	Permissions                        string       `json:"permissions"`
	SlowModeDelay                      int          `json:"slow_mode_delay"`
	MessageAutoDeleteTime              int          `json:"message_auto_delete_time"`
	HasAggressiveAntiSpamEnabled       bool         `json:"has_aggressive_anti_spam_enabled"`
	HasHiddenMembers                   bool         `json:"has_hidden_members"`
	HasProtectedContent                bool         `json:"has_protected_content"`
	StickerSetName                     string       `json:"sticker_set_name"`
	CanSetStickerSet                   bool         `json:"can_set_sticker_set"`
	LinkedChatId                       int          `json:"linked_chat_id"`
	Location                           ChatLocation `json:"location"`
}

//https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

//https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendAudios         bool `json:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
	CanManageTopics       bool `json:"can_manage_topics"`
}

//https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}
