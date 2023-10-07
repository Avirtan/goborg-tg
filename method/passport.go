package method

import (
	"context"
	"encoding/json"
	"log/slog"

	method_dto "github.com/Avirtan/TGoBot/dto/method"
	"github.com/Avirtan/TGoBot/request"
)

func SetPassportDataErrors(ctx context.Context, spde method_dto.SetPassportDataErrors) (*bool, error) {
	data, err := json.Marshal(spde)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setPassportDataErrors", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}
