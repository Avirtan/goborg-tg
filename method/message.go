package method

import (
	method_dto "TGoBot/dto/method"
	"TGoBot/request"
	"context"
	"encoding/json"
	"log/slog"
)

// https://core.telegram.org/bots/api#sendmessage
func SendMessage(ctx context.Context, msg method_dto.SendMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendMessage", data)
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

// https://core.telegram.org/bots/api#forwardmessage
func ForwardMessage(ctx context.Context, msg method_dto.ForwardMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/forwardMessage", data)
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

// https://core.telegram.org/bots/api#copymessage
func CopyMessage(ctx context.Context, msg method_dto.CopyMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/copyMessage", data)
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
