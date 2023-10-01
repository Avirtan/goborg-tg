package method

import (
	method_dto "TGoBot/dto/method"
	"TGoBot/request"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
)

func SendMessage(ctx context.Context, msg method_dto.SendMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
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
