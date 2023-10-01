package method

import (
	"TGoBot/request"
	"context"
	"log/slog"
)

func GetMe(ctx context.Context) error {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/answerInlineQuery")
	if err != nil {
		return err
	}
	responseJson, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}
	slog.Info("info", "response", responseJson)
	return nil
}
