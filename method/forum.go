package method

import (
	"context"
	"log/slog"

	forum_dto "github.com/Avirtan/TGoBot/dto/forum"
	stickers_dto "github.com/Avirtan/TGoBot/dto/stickers"
	"github.com/Avirtan/TGoBot/request"
)

func GetForumTopicIconStickers(ctx context.Context) (*[]stickers_dto.Sticker, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/getForumTopicIconStickers")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[[]stickers_dto.Sticker](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

func CreateForumTopic[ID int64 | string](ctx context.Context, name string) (*forum_dto.ForumTopic, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/createForumTopic")
	if err != nil {
		return nil, err
	}
	responseJson, err := request.ResponseHandlerToType[forum_dto.ForumTopic](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

func EditForumTopic[ID int64 | string](ctx context.Context, messageThreadId int64) (*bool, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/editForumTopic")
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

func CloseForumTopic[ID int64 | string](ctx context.Context, messageThreadId int64) (*bool, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/closeForumTopic")
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

func ReopenForumTopic[ID int64 | string](ctx context.Context, messageThreadId int64) (*bool, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/reopenForumTopic")
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

func DeleteForumTopic[ID int64 | string](ctx context.Context, messageThreadId int64) (*bool, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/deleteForumTopic")
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
