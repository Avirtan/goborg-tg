package user_dto

import utils_dto "TGoBot/dto/utils"

//https://core.telegram.org/bots/api#user
type User struct {
	Id                      int64  `json:"id"`
	FirstName               string `json:"first_name,omitempty"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsBot                   bool   `json:"is_bot"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

//https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int64                   `json:"total_count"`
	Photos     [][]utils_dto.PhotoSize `json:"photos"`
}
