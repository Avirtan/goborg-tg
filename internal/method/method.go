package method

import (
	"TGoBot/internal/dto"
	"TGoBot/internal/request"
	"context"
	"encoding/json"
	"fmt"
	"io"
)

const baseUrl = "https://api.telegram.org/bot%s"

type MethodHandler struct {
	Token string
}

func New(token string) *MethodHandler {
	return &MethodHandler{
		Token: token,
	}
}

func (m *MethodHandler) GetUrl() string {
	return fmt.Sprintf(baseUrl, m.Token)
}

func (m *MethodHandler) GetUpdates(ctx context.Context, offset uint64, limit int, timeout int) (*dto.UpdateResponse, error) {
	dataRequest := dto.UpdateRequest{
		Offset:  offset,
		Limit:   limit,
		Timeout: timeout,
	}
	data, err := json.Marshal(dataRequest)
	if err != nil {
		return nil, err
	}
	resp, err := request.RequestWithContext(ctx, request.Get, m.GetUrl()+"/getUpdates", data)
	if err != nil {
		return nil, err
	}
	response := dto.UpdateResponse{}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &response, err
}

func (m *MethodHandler) SetMyCommands(ctx context.Context, command *dto.BotCommand) error {
	botCommandRequest := dto.BotCommandRequst{
		Command:     command.Command,
		Description: command.Description,
	}
	dataRequest := dto.SetCommandRequest{
		Commands: []dto.BotCommandRequst{botCommandRequest},
		// LanguageCode: "ru",
	}
	data, err := json.Marshal(dataRequest)
	if err != nil {
		return err
	}
	resp, err := request.RequestWithContext(ctx, request.Get, m.GetUrl()+"/setMyCommands", data)
	if err != nil {
		return err
	}
	// var response interface{}
	// body, _ := io.ReadAll(resp.Body)
	// err = json.Unmarshal(body, &response)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(response)
	resp.Body.Close()
	return nil
}

func (m *MethodHandler) DeleteMyCommands() error {
	resp, err := request.Request(request.Get, m.GetUrl()+"/deleteMyCommands", []byte{})
	if err != nil {
		return err
	}
	// var response interface{}
	// body, _ := io.ReadAll(resp.Body)
	// err = json.Unmarshal(body, &response)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(response)
	resp.Body.Close()
	return nil
}

func (m *MethodHandler) GetMyCommands() error {
	resp, err := request.Request(request.Get, m.GetUrl()+"/getMyCommands", []byte{})
	if err != nil {
		return err
	}
	var response interface{}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	fmt.Println(response)
	resp.Body.Close()
	return nil
}

func (m *MethodHandler) SendMessage(msg dto.SendMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := request.Request(request.Get, m.GetUrl()+"/sendMessage", data)
	if err != nil {
		return err
	}
	var response interface{}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}
