package method

import (
	"context"
	"encoding/json"
	"log/slog"

	method_dto "github.com/Avirtan/TGoBot/dto/method"
	user_dto "github.com/Avirtan/TGoBot/dto/user"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#getuserprofilephotos
func GetUserProfilePhotos(ctx context.Context, user method_dto.GetUserProfilePhotos) (*user_dto.UserProfilePhotos, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getUserProfilePhotos", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[user_dto.UserProfilePhotos](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}
