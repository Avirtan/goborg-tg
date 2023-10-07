package method

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/Avirtan/TGoBot/dto"
	method_dto "github.com/Avirtan/TGoBot/dto/method"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#setpassportdataerrors
func SetPassportDataErrors(ctx context.Context, data method_dto.SetPassportDataErrors) (*dto.Response, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setPassportDataErrors", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}
