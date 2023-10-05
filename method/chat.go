package method

import (
	"context"
	"log/slog"

	message_dto "github.com/Avirtan/TGoBot/dto/message"
	"github.com/Avirtan/TGoBot/request"
)

func GetChat[ID int64 | string](ctx context.Context) (*message_dto.Chat, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/getChat")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[message_dto.Chat](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

func GetChatAdministrators[ID int64 | string](ctx context.Context) (*message_dto.ChatMember, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/getChatAdministrators")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[message_dto.ChatMember](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

func GetChatMemberCount[ID int64 | string](ctx context.Context) (*int64, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/getChatMemberCount")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[int64](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

func GetChatMember[ID int64 | string](ctx context.Context, userId int64) (*message_dto.ChatMember, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/getChatMember")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[message_dto.ChatMember](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

func SetChatStickerSet[ID int64 | string](ctx context.Context, stickerSetName string) (*bool, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/setChatStickerSet")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}
