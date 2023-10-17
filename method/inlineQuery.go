package method

import (
	"context"
	"encoding/json"
	"log/slog"

	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	"github.com/Avirtan/goborg-tg/request"
)

// https://core.telegram.org/bots/api#answerinlinequery
func AnswerInlineQuery(ctx context.Context, data method_dto.AnswerInlineQuery) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/answerInlineQuery", marshalBytes)
	if err != nil {
		return err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}

	slog.Debug("info", "response", responseData)

	return nil
}

// https://core.telegram.org/bots/api#answercallbackquery
func AnswerCallbackQuery(ctx context.Context, data method_dto.AnswerCallbackQuery) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/answerCallbackQuery", marshalBytes)
	if err != nil {
		return err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}

	slog.Debug("info", "response", responseData)

	return nil
}
