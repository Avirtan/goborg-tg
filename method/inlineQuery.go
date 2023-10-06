package method

import (
	"context"
	"encoding/json"
	"log/slog"

	method_dto "github.com/Avirtan/TGoBot/dto/method"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#answerinlinequery
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

// https://core.telegram.org/bots/api#answercallbackquery
func AnswerCallbackQuery(ctx context.Context, msg method_dto.AnswerCallbackQuery) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/answerCallbackQuery", data)
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
