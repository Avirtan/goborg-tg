package method_dto

// https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotos struct {
	UserId uint64 `json:"user_id"`
	Offset uint16 `json:"offset,omitempty"`
	Limit  uint8  `json:"limit,omitempty"`
}

func NewGetUserProfilePhotos(userId uint64) *GetUserProfilePhotos {
	return &GetUserProfilePhotos{
		UserId: userId,
	}
}

func (g *GetUserProfilePhotos) SetOffset(offset uint16) *GetUserProfilePhotos {
	g.Offset = offset
	return g
}

func (g *GetUserProfilePhotos) SetLimit(limit uint8) *GetUserProfilePhotos {
	g.Limit = limit
	return g
}
