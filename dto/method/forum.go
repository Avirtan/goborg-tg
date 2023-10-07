package method_dto

// https://core.telegram.org/bots/api#createforumtopic
type CreateForumTopic[ID int64 | string] struct {
	ChatID            ID     `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

func NewCreateForumTopic[ID int64 | string](chatID ID, name string) *CreateForumTopic[ID] {
	return &CreateForumTopic[ID]{
		ChatID: chatID,
		Name:   name,
	}
}

func (ft *CreateForumTopic[ID]) SetIconColor(iconColor int) *CreateForumTopic[ID] {
	ft.IconColor = iconColor
	return ft
}

func (ft *CreateForumTopic[ID]) SetIconCustomEmojiID(iconCustomEmojiID string) *CreateForumTopic[ID] {
	ft.IconCustomEmojiID = iconCustomEmojiID
	return ft
}

// https://core.telegram.org/bots/api#editforumtopic
type EditForumTopic[ID int64 | string] struct {
	ChatID            ID     `json:"chat_id"`
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

func NewEditForumTopic[ID int64 | string](chatID ID, messageThreadID int) *EditForumTopic[ID] {
	return &EditForumTopic[ID]{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
}

func (e *EditForumTopic[ID]) SetName(name string) *EditForumTopic[ID] {
	e.Name = name
	return e
}

func (e *EditForumTopic[ID]) SetIconCustomEmojiID(iconCustomEmojiID string) *EditForumTopic[ID] {
	e.IconCustomEmojiID = iconCustomEmojiID
	return e
}

// https://core.telegram.org/bots/api#closeforumtopic
type CloseForumTopic[ID int64 | string] struct {
	ChatID          ID  `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

func NewCloseForumTopic[ID int64 | string](chatID ID, messageThreadID int) *CloseForumTopic[ID] {
	return &CloseForumTopic[ID]{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
}

// https://core.telegram.org/bots/api#reopenforumtopic
type ReopenForumTopic[ID int64 | string] struct {
	ChatID          ID  `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

func NewReopenForumTopic[ID int64 | string](chatID ID, messageThreadID int) *ReopenForumTopic[ID] {
	return &ReopenForumTopic[ID]{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
}

// https://core.telegram.org/bots/api#deleteforumtopic
type DeleteForumTopic[ID int64 | string] struct {
	ChatID          ID  `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

func NewDeleteForumTopic[ID int64 | string](chatID ID, messageThreadID int) *DeleteForumTopic[ID] {
	return &DeleteForumTopic[ID]{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
}

// https://core.telegram.org/bots/api#unpinallforumtopicmessages
type UnpinAllForumTopicMessages[ID int64 | string] struct {
	ChatID          ID  `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

func NewUnpinAllForumTopicMessages[ID int64 | string](chatID ID, messageThreadID int) *UnpinAllForumTopicMessages[ID] {
	return &UnpinAllForumTopicMessages[ID]{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
}

// https://core.telegram.org/bots/api#editgeneralforumtopic
type EditGeneralForumTopic[ID int64 | string] struct {
	ChatID ID     `json:"chat_id"`
	Name   string `json:"name"`
}

func NewEditGeneralForumTopic[ID int64 | string](chatID ID, name string) *EditGeneralForumTopic[ID] {
	return &EditGeneralForumTopic[ID]{
		ChatID: chatID,
		Name:   name,
	}
}
