package dto

// https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}
