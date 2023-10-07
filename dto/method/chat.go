package method_dto

import message_dto "github.com/Avirtan/TGoBot/dto/message"

// ChatID
type ChatID[ID int64 | string] struct {
	ChatID ID `json:"chat_id"`
}

func NewChatID[ID int64 | string](chatID ID) *ChatID[ID] {
	return &ChatID[ID]{
		ChatID: chatID,
	}
}

// https://core.telegram.org/bots/api#banchatmember
type BanChatMember[ID int64 | string] struct {
	ChatID         ID   `json:"chat_id"`
	UserID         int  `json:"user_id"`
	UntilDate      int  `json:"until_date,omitempty"`
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

func NewBanChatMember[ID int64 | string](chatID ID, userID int) *BanChatMember[ID] {
	return &BanChatMember[ID]{
		ChatID: chatID,
		UserID: userID,
	}
}

func (bcm *BanChatMember[ID]) SetUntilDate(untilDate int) *BanChatMember[ID] {
	bcm.UntilDate = untilDate
	return bcm
}

func (bcm *BanChatMember[ID]) SetRevokeMessages(revoke bool) *BanChatMember[ID] {
	bcm.RevokeMessages = revoke
	return bcm
}

// https://core.telegram.org/bots/api#unbanchatmember
type UnbanChatMember[ID int64 | string] struct {
	ChatID       ID   `json:"chat_id"`
	UserID       int  `json:"user_id"`
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

func NewUnbanChatMember[ID int64 | string](chatID ID, userID int) *UnbanChatMember[ID] {
	return &UnbanChatMember[ID]{
		ChatID: chatID,
		UserID: userID,
	}
}

func (ucm *UnbanChatMember[ID]) SetOnlyIfBanned(onlyIfBanned bool) *UnbanChatMember[ID] {
	ucm.OnlyIfBanned = onlyIfBanned
	return ucm
}

// https://core.telegram.org/bots/api#restrictchatmember
type RestrictChatMember[ID int64 | string] struct {
	ChatID                        ID                          `json:"chat_id"`
	UserID                        int                         `json:"user_id"`
	Permissions                   message_dto.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                        `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int                         `json:"until_date,omitempty"`
}

func NewRestrictChatMember[ID int64 | string](chatID ID, userID int, permissions message_dto.ChatPermissions) *RestrictChatMember[ID] {
	return &RestrictChatMember[ID]{
		ChatID:      chatID,
		UserID:      userID,
		Permissions: permissions,
	}
}

func (rcm *RestrictChatMember[ID]) SetUseIndependentChatPermissions(useIndependent bool) *RestrictChatMember[ID] {
	rcm.UseIndependentChatPermissions = useIndependent
	return rcm
}

func (rcm *RestrictChatMember[ID]) SetUntilDate(untilDate int) *RestrictChatMember[ID] {
	rcm.UntilDate = untilDate
	return rcm
}

// https://core.telegram.org/bots/api#promotechatmember
type PromoteChatMember[ID int64 | string] struct {
	ChatID              ID   `json:"chat_id"`
	UserID              int  `json:"user_id"`
	IsAnonymous         bool `json:"is_anonymous,omitempty"`
	CanManageChat       bool `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool `json:"can_change_info,omitempty"`
	CanInviteUsers      bool `json:"can_invite_users,omitempty"`
	CanPostMessages     bool `json:"can_post_messages,omitempty"`
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`
	CanPostStories      bool `json:"can_post_stories,omitempty"`
	CanEditStories      bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`
}

func NewPromoteChatMember[ID int64 | string](chatID ID, userID int) *PromoteChatMember[ID] {
	return &PromoteChatMember[ID]{
		ChatID: chatID,
		UserID: userID,
	}
}

func (pcm *PromoteChatMember[ID]) SetIsAnonymous(isAnonymous bool) *PromoteChatMember[ID] {
	pcm.IsAnonymous = isAnonymous
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanManageChat(canManageChat bool) *PromoteChatMember[ID] {
	pcm.CanManageChat = canManageChat
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanDeleteMessages(canDeleteMessages bool) *PromoteChatMember[ID] {
	pcm.CanDeleteMessages = canDeleteMessages
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanManageVideoChats(canManageVideoChats bool) *PromoteChatMember[ID] {
	pcm.CanManageVideoChats = canManageVideoChats
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanRestrictMembers(canRestrictMembers bool) *PromoteChatMember[ID] {
	pcm.CanRestrictMembers = canRestrictMembers
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanPromoteMembers(canPromoteMembers bool) *PromoteChatMember[ID] {
	pcm.CanPromoteMembers = canPromoteMembers
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanChangeInfo(canChangeInfo bool) *PromoteChatMember[ID] {
	pcm.CanChangeInfo = canChangeInfo
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanInviteUsers(canInviteUsers bool) *PromoteChatMember[ID] {
	pcm.CanInviteUsers = canInviteUsers
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanPostMessages(canPostMessages bool) *PromoteChatMember[ID] {
	pcm.CanPostMessages = canPostMessages
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanEditMessages(canEditMessages bool) *PromoteChatMember[ID] {
	pcm.CanEditMessages = canEditMessages
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanPinMessages(canPinMessages bool) *PromoteChatMember[ID] {
	pcm.CanPinMessages = canPinMessages
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanPostStories(canPostStories bool) *PromoteChatMember[ID] {
	pcm.CanPostStories = canPostStories
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanEditStories(canEditStories bool) *PromoteChatMember[ID] {
	pcm.CanEditStories = canEditStories
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanDeleteStories(canDeleteStories bool) *PromoteChatMember[ID] {
	pcm.CanDeleteStories = canDeleteStories
	return pcm
}

func (pcm *PromoteChatMember[ID]) SetCanManageTopics(canManageTopics bool) *PromoteChatMember[ID] {
	pcm.CanManageTopics = canManageTopics
	return pcm
}

// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
type SetChatAdministratorCustomTitle[ID int64 | string] struct {
	ChatID      ID     `json:"chat_id"`
	UserID      int    `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

func NewSetChatAdministratorCustomTitle[ID int64 | string](chatID ID, userID int, customTitle string) *SetChatAdministratorCustomTitle[ID] {
	return &SetChatAdministratorCustomTitle[ID]{
		ChatID:      chatID,
		UserID:      userID,
		CustomTitle: customTitle,
	}
}

// https://core.telegram.org/bots/api#banchatsenderchat
type BanChatSenderChat[ID int64 | string] struct {
	ChatID       ID  `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

func NewBanChatSenderChat[ID int64 | string](chatID ID, senderChatID int) *BanChatSenderChat[ID] {
	return &BanChatSenderChat[ID]{
		ChatID:       chatID,
		SenderChatID: senderChatID,
	}
}

// https://core.telegram.org/bots/api#unbanchatsenderchat
type UnbanChatSenderChat[ID int64 | string] struct {
	ChatID       ID  `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

func NewUnbanChatSenderChat[ID int64 | string](chatID ID, senderChatID int) *UnbanChatSenderChat[ID] {
	return &UnbanChatSenderChat[ID]{
		ChatID:       chatID,
		SenderChatID: senderChatID,
	}
}

// https://core.telegram.org/bots/api#setchatpermissions
type SetChatPermissions[ID int64 | string] struct {
	ChatID                        ID                          `json:"chat_id"`
	Permissions                   message_dto.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                        `json:"use_independent_chat_permissions,omitempty"`
}

func NewSetChatPermissions[ID int64 | string](chatID ID, permissions message_dto.ChatPermissions) *SetChatPermissions[ID] {
	return &SetChatPermissions[ID]{
		ChatID:      chatID,
		Permissions: permissions,
	}
}

func (scp *SetChatPermissions[ID]) SetUseIndependentChatPermissions(useIndependent bool) *SetChatPermissions[ID] {
	scp.UseIndependentChatPermissions = useIndependent
	return scp
}

// https://core.telegram.org/bots/api#createchatinvitelink
type CreateChatInviteLink[ID int64 | string] struct {
	ChatID             ID     `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

func NewCreateChatInviteLink[ID int64 | string](chatID ID) *CreateChatInviteLink[ID] {
	return &CreateChatInviteLink[ID]{
		ChatID: chatID,
	}
}

func (c *CreateChatInviteLink[ID]) SetName(name string) *CreateChatInviteLink[ID] {
	c.Name = name
	return c
}

func (c *CreateChatInviteLink[ID]) SetExpireDate(expireDate int64) *CreateChatInviteLink[ID] {
	c.ExpireDate = expireDate
	return c
}

func (c *CreateChatInviteLink[ID]) SetMemberLimit(memberLimit int) *CreateChatInviteLink[ID] {
	c.MemberLimit = memberLimit
	return c
}

func (c *CreateChatInviteLink[ID]) SetCreatesJoinRequest(createsJoinRequest bool) *CreateChatInviteLink[ID] {
	c.CreatesJoinRequest = createsJoinRequest
	return c
}

// https://core.telegram.org/bots/api#editchatinvitelink
type EditChatInviteLink[ID int64 | string] struct {
	ChatID             ID     `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

func NewEditChatInviteLink[ID int64 | string](chatID ID, inviteLink string) *EditChatInviteLink[ID] {
	return &EditChatInviteLink[ID]{
		ChatID:     chatID,
		InviteLink: inviteLink,
	}
}

func (e *EditChatInviteLink[ID]) SetName(name string) *EditChatInviteLink[ID] {
	e.Name = name
	return e
}

func (e *EditChatInviteLink[ID]) SetExpireDate(expireDate int64) *EditChatInviteLink[ID] {
	e.ExpireDate = expireDate
	return e
}

func (e *EditChatInviteLink[ID]) SetMemberLimit(memberLimit int) *EditChatInviteLink[ID] {
	e.MemberLimit = memberLimit
	return e
}

func (e *EditChatInviteLink[ID]) SetCreatesJoinRequest(createsJoinRequest bool) *EditChatInviteLink[ID] {
	e.CreatesJoinRequest = createsJoinRequest
	return e
}

// https://core.telegram.org/bots/api#revokechatinvitelink
type RevokeChatInviteLink[ID int64 | string] struct {
	ChatID     ID     `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

func NewRevokeChatInviteLink[ID int64 | string](chatID ID, inviteLink string) *RevokeChatInviteLink[ID] {
	return &RevokeChatInviteLink[ID]{
		ChatID:     chatID,
		InviteLink: inviteLink,
	}
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
type ApproveChatJoinRequest[ID int64 | string] struct {
	ChatID ID  `json:"chat_id"`
	UserID int `json:"user_id"`
}

func NewApproveChatJoinRequest[ID int64 | string](chatID ID, userID int) *ApproveChatJoinRequest[ID] {
	return &ApproveChatJoinRequest[ID]{
		ChatID: chatID,
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#declinechatjoinrequest
type DeclineChatJoinRequest[ID int64 | string] struct {
	ChatID ID  `json:"chat_id"`
	UserID int `json:"user_id"`
}

func NewDeclineChatJoinRequest[ID int64 | string](chatID ID, userID int) *DeclineChatJoinRequest[ID] {
	return &DeclineChatJoinRequest[ID]{
		ChatID: chatID,
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#setchatphoto
type SetChatPhoto[ID int64 | string] struct {
	ChatID ID     `json:"chat_id"`
	Photo  string `json:"photo"`
}

func NewSetChatPhoto[ID int64 | string](chatID ID, photo string) *SetChatPhoto[ID] {
	return &SetChatPhoto[ID]{
		ChatID: chatID,
		Photo:  photo,
	}
}

func (scp *SetChatPhoto[ID]) SetPhoto(photo string) *SetChatPhoto[ID] {
	scp.Photo = photo
	return scp
}

// https://core.telegram.org/bots/api#setchattitle
type SetChatTitle[ID int64 | string] struct {
	ChatID ID     `json:"chat_id"`
	Title  string `json:"title"`
}

func NewSetChatTitle[ID int64 | string](chatID ID, title string) *SetChatTitle[ID] {
	return &SetChatTitle[ID]{
		ChatID: chatID,
		Title:  title,
	}
}

func (sct *SetChatTitle[ID]) SetTitle(title string) *SetChatTitle[ID] {
	sct.Title = title
	return sct
}

// https://core.telegram.org/bots/api#setchatdescription
type SetChatDescription[ID int64 | string] struct {
	ChatID      ID     `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

func NewSetChatDescription[ID int64 | string](chatID ID, description string) *SetChatDescription[ID] {
	return &SetChatDescription[ID]{
		ChatID:      chatID,
		Description: description,
	}
}

func (scd *SetChatDescription[ID]) SetDescription(description string) *SetChatDescription[ID] {
	scd.Description = description
	return scd
}

// https://core.telegram.org/bots/api#pinchatmessage
type PinChatMessage[ID int64 | string] struct {
	ChatID              ID   `json:"chat_id"`
	MessageID           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
}

func NewPinChatMessage[ID int64 | string](chatID ID, messageID int) *PinChatMessage[ID] {
	return &PinChatMessage[ID]{
		ChatID:    chatID,
		MessageID: messageID,
	}
}

func (pcm *PinChatMessage[ID]) SetDisableNotification(disable bool) *PinChatMessage[ID] {
	pcm.DisableNotification = disable
	return pcm
}

// https://core.telegram.org/bots/api#unpinchatmessage
type UnpinChatMessage[ID int64 | string] struct {
	ChatID    ID  `json:"chat_id"`
	MessageID int `json:"message_id,omitempty"`
}

func NewUnpinChatMessage[ID int64 | string](chatID ID) *UnpinChatMessage[ID] {
	return &UnpinChatMessage[ID]{
		ChatID: chatID,
	}
}

func (ucm *UnpinChatMessage[ID]) SetMessageID(messageID int) *UnpinChatMessage[ID] {
	ucm.MessageID = messageID
	return ucm
}

// https://core.telegram.org/bots/api#getchatmember
type GetChatMember[ID int64 | string] struct {
	ChatID ID  `json:"chat_id"`
	UserID int `json:"user_id"`
}

func NewGetChatMember[ID int64 | string](chatID ID, userID int) *GetChatMember[ID] {
	return &GetChatMember[ID]{
		ChatID: chatID,
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#setchatstickerset
type SetChatStickerSet[ID int64 | string] struct {
	ChatID         ID     `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

func NewSetChatStickerSet[ID int64 | string](chatID ID, stickerSetName string) *SetChatStickerSet[ID] {
	return &SetChatStickerSet[ID]{
		ChatID:         chatID,
		StickerSetName: stickerSetName,
	}
}
