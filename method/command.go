package method

import (
	"context"
	"encoding/json"
	"log/slog"

	command_dto "github.com/Avirtan/goborg-tg/dto/command"
	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	"github.com/Avirtan/goborg-tg/request"
)

// https://core.telegram.org/bots/api#setmycommands
func SetMyCommands(ctx context.Context, data []*command_dto.BotCommand) error {
	dataRequest := method_dto.SetCommandRequest{
		Commands: data,
		// LanguageCode: "ru",
	}

	marshalBytes, err := json.Marshal(dataRequest)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyCommands", marshalBytes)
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

// https://core.telegram.org/bots/api#deletemycommands
func DeleteMyCommands(ctx context.Context) error {
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteMyCommands", []byte{})
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

// https://core.telegram.org/bots/api#getmycommands
func GetMyCommands(ctx context.Context) error {
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyCommands", []byte{})
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
