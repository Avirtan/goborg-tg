package method_dto

// https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotos struct {
	UserId uint64 `json:"user_id"`
	Offset uint16 `json:"offset,omitempty"`
	Limit  uint8  `json:"limit,omitempty"`
}
