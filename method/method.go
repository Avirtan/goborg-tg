package method

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	update_dto "github.com/Avirtan/TGoBot/dto/update"
	"github.com/Avirtan/TGoBot/request"
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
