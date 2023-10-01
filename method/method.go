package method

import (
	command_dto "TGoBot/dto/command"
	method_dto "TGoBot/dto/method"
	update_dto "TGoBot/dto/update"
	"TGoBot/request"
	"log/slog"

	"context"
	"encoding/json"
	"fmt"
	"io"
)

const baseUrl = "https://api.telegram.org/bot%s"

var tokenNew string

func New(token string) {
	tokenNew = token
}

func GetUrl() string {
	return fmt.Sprintf(baseUrl, tokenNew)
}

func GetUpdates(ctx context.Context, offset uint64, limit int, timeout int) (*update_dto.UpdateResponse, error) {
	dataRequest := update_dto.UpdateRequest{
		Offset:  offset,
		Limit:   limit,
		Timeout: timeout,
	}
	data, err := json.Marshal(dataRequest)
	if err != nil {
		return nil, err
	}
	resp, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/getUpdates", data)
	if err != nil {
		return nil, err
	}
	response := update_dto.UpdateResponse{}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &response, err
}

func SetMyCommands(ctx context.Context, commands []*command_dto.BotCommand) error {
	dataRequest := method_dto.SetCommandRequest{
		Commands: commands,
		// LanguageCode: "ru",
	}
	data, err := json.Marshal(dataRequest)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/setMyCommands", data)
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

func DeleteMyCommands() error {
	response, err := request.Request(request.Get, GetUrl()+"/deleteMyCommands", []byte{})
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

func GetMyCommands() error {
	response, err := request.Request(request.Get, GetUrl()+"/getMyCommands", []byte{})
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

func SendMessage(msg method_dto.SendMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.Request(request.Get, GetUrl()+"/sendMessage", data)
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

func AnswerInlineQuery(msg method_dto.AnswerInlineQuery) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.Request(request.Get, GetUrl()+"/answerInlineQuery", data)
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
