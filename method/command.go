package method

import (
	command_dto "TGoBot/dto/command"
	method_dto "TGoBot/dto/method"
	"TGoBot/request"
	"context"
	"encoding/json"
	"log/slog"
)

func SetMyCommands(ctx context.Context, commands []*command_dto.BotCommand) error {
	dataRequest := method_dto.SetCommandRequest{
		Commands: commands,
		// LanguageCode: "ru",
	}
	data, err := json.Marshal(dataRequest)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyCommands", data)
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

func DeleteMyCommands(ctx context.Context) error {
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteMyCommands", []byte{})
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

func GetMyCommands(ctx context.Context) error {
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyCommands", []byte{})
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
