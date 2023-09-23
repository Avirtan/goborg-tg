package dto

//https://core.telegram.org/bots/api#user
type User struct {
	Id                      int64  `json:"id"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsBot                   bool   `json:"is_bot"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}
