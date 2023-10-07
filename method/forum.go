package method

import (
	"context"
	"encoding/json"
	"log/slog"

	forum_dto "github.com/Avirtan/TGoBot/dto/forum"
	method_dto "github.com/Avirtan/TGoBot/dto/method"
	stickers_dto "github.com/Avirtan/TGoBot/dto/stickers"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#getforumtopiciconstickers
func GetForumTopicIconStickers(ctx context.Context) (*[]stickers_dto.Sticker, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/getForumTopicIconStickers")
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[[]stickers_dto.Sticker](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#createforumtopic
func CreateForumTopic[ID int64 | string](ctx context.Context, data method_dto.CreateForumTopic[ID]) (*forum_dto.ForumTopic, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/createForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[forum_dto.ForumTopic](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#editforumtopic
func EditForumTopic[ID int64 | string](ctx context.Context, data method_dto.EditForumTopic[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#closeforumtopic
func CloseForumTopic[ID int64 | string](ctx context.Context, data method_dto.CloseForumTopic[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/closeForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#reopenforumtopic
func ReopenForumTopic[ID int64 | string](ctx context.Context, data method_dto.ReopenForumTopic[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/reopenForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#deleteforumtopic
func DeleteForumTopic[ID int64 | string](ctx context.Context, data method_dto.DeleteForumTopic[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func UnpinAllForumTopicMessages[ID int64 | string](ctx context.Context, data method_dto.UnpinAllForumTopicMessages[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinAllForumTopicMessages", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#editgeneralforumtopic
func EditGeneralForumTopic[ID int64 | string](ctx context.Context, data method_dto.EditGeneralForumTopic[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editGeneralForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#closegeneralforumtopic
func CloseGeneralForumTopic[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/closeGeneralForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#reopengeneralforumtopic
func ReopenGeneralForumTopic[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/reopenGeneralForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#hidegeneralforumtopic
func HideGeneralForumTopic[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/hideGeneralForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func UnhideGeneralForumTopic[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unhideGeneralForumTopic", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func UnpinAllGeneralForumTopicMessages[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinAllGeneralForumTopicMessages", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bool](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}
