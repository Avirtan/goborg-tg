package method

import (
	method_dto "TGoBot/dto/method"
	"TGoBot/request"
	"context"
	"encoding/json"
	"log/slog"
)

func AnswerInlineQuery(ctx context.Context, msg method_dto.AnswerInlineQuery) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/answerInlineQuery", data)
	if err != nil {
		return err
	}
	responseJson, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}
	slog.Debug("info", "response", responseJson)
	return nil
}
