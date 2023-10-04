package message_dto

import (
	user_dto "github.com/Avirtan/TGoBot/dto/user"
	utils_dto "github.com/Avirtan/TGoBot/dto/utils"
)

// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id                                 int64        `json:"id"`
	Type                               string       `json:"type"`
	Title                              string       `json:"title"`
	Username                           string       `json:"username"`
	FirstName                          string       `json:"first_name"`
	LastName                           string       `json:"last_name"`
	EmojiStatusCustomEmojiId           string       `json:"added_to_attachment_menu"`
	Bio                                string       `json:"bio"`
	Description                        string       `json:"description"`
	InviteLink                         string       `json:"invite_link"`
	Permissions                        string       `json:"permissions"`
	StickerSetName                     string       `json:"sticker_set_name"`
	Photo                              ChatPhoto    `json:"photo"`
	ActiveUsernames                    []string     `json:"active_usernames"`
	EmojiStatusExpirationDate          int          `json:"emoji_status_expiration_date"`
	SlowModeDelay                      int          `json:"slow_mode_delay"`
	MessageAutoDeleteTime              int          `json:"message_auto_delete_time"`
	HasAggressiveAntiSpamEnabled       bool         `json:"has_aggressive_anti_spam_enabled"`
	HasHiddenMembers                   bool         `json:"has_hidden_members"`
	HasProtectedContent                bool         `json:"has_protected_content"`
	CanSetStickerSet                   bool         `json:"can_set_sticker_set"`
	IsForum                            bool         `json:"is_forum"`
	HasPrivateForwards                 bool         `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages bool         `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 bool         `json:"join_to_send_messages"`
	JoinByRequest                      bool         `json:"join_by_request"`
	LinkedChatId                       int          `json:"linked_chat_id"`
	Location                           ChatLocation `json:"location"`
	PinnedMessage                      *Message     `json:"pinned_message"`
}

// https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

// https://core.telegram.org/bots/api#chatpermissions
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

// https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	Location utils_dto.Location `json:"location"`
	Address  string             `json:"address"`
}

// https://core.telegram.org/bots/api#chatinvitelink
type ChatInviteLink struct {
	InviteLink              string        `json:"invite_link"`
	Creator                 user_dto.User `json:"creator"`
	CreatesJoinRequest      bool          `json:"creates_join_request"`
	IsPrimary               bool          `json:"is_primary"`
	IsRevoked               bool          `json:"is_revoked"`
	Name                    string        `json:"name,omitempty"`
	ExpireDate              int64         `json:"expire_date,omitempty"`
	MemberLimit             int           `json:"member_limit,omitempty"`
	PendingJoinRequestCount int           `json:"pending_join_request_count,omitempty"`
}

// https://core.telegram.org/bots/api#chatadministratorrights
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanPostMessages     bool `json:"can_post_messages,omitempty"`
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`
	CanPostStories      bool `json:"can_post_stories,omitempty"`
	CanEditStories      bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`
}

// https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	Chat       Chat           `json:"chat"`
	From       user_dto.User  `json:"from"`
	UserChatId int64          `json:"user_chat_id"`
	Date       int64          `json:"date"`
	Bio        string         `json:"bio"`
	InviteLink ChatInviteLink `json:"invite_link"`
}
