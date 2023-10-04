package method

import (
	"context"
	"log/slog"

	user_dto "github.com/Avirtan/TGoBot/dto/user"
	"github.com/Avirtan/TGoBot/request"
)

func GetMe(ctx context.Context) (*user_dto.User, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/GetMe")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[user_dto.User](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}
