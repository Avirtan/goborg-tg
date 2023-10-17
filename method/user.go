package method

import (
	"context"
	"encoding/json"
	"log/slog"

	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	user_dto "github.com/Avirtan/goborg-tg/dto/user"
	"github.com/Avirtan/goborg-tg/request"
)

// https://core.telegram.org/bots/api#getuserprofilephotos
func GetUserProfilePhotos(ctx context.Context, data method_dto.GetUserProfilePhotos) (*user_dto.UserProfilePhotos, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getUserProfilePhotos", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[user_dto.UserProfilePhotos](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}
