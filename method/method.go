package method

import (
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
	resp, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getUpdates", data)
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
