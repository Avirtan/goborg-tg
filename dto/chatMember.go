package dto

// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	ChatMemberOwner
	ChatMemberAdministrator
	ChatMemberMember
	ChatMemberRestricted
	ChatMemberLeft
	ChatMemberBanned
}

// https://core.telegram.org/bots/api#chatmemberowner
type ChatMemberOwner struct {
	Status      string `json:"status"`
	User        User   `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title"`
}

// https://core.telegram.org/bots/api#chatmemberadministrator
type ChatMemberAdministrator struct {
	Status              string `json:"status"`
	User                User   `json:"user"`
	CanBeEdited         bool   `json:"can_be_edited"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanDeleteMessages   bool   `json:"ican_delete_messagess_anonymous"`
	CanManageVideoChats bool   `json:"can_manage_video_chats"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanChangeInfo       bool   `json:"can_change_info"`
	CanInviteUsers      bool   `json:"can_invite_users"`
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`
	CanPostStories      bool   `json:"can_post_stories,omitempty"`
	CanEditStories      bool   `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool   `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`
	CustomTitle         string `json:"custom_title,omitempty"`
}

// https://core.telegram.org/bots/api#chatmembermember
type ChatMemberMember struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

// https://core.telegram.org/bots/api#chatmemberrestricted
type ChatMemberRestricted struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	UntilDate             int    `json:"until_date"`
}

// https://core.telegram.org/bots/api#chatmemberleft
type ChatMemberLeft struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

// https://core.telegram.org/bots/api#chatmemberbanned
type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate int    `json:"until_date"`
}

// https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	Chat                    Chat       `json:"chat"`
	From                    User       `json:"from"`
	Date                    int        `json:"date"`
	OldChatMember           ChatMember `json:"old_chat_member"`
	NewChatMember           ChatMember `json:"new_chat_member"`
	InviteLink              ChatMember `json:"invite_link,omitempty"`
	ViaChatFolderInviteLink bool       `json:"via_chat_folder_invite_link,omitempty"`
}
