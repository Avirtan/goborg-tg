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

	responseJson, err := request.ResponseHandlerToType[[]stickers_dto.Sticker](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#createforumtopic
func CreateForumTopic[ID int64 | string](ctx context.Context, cft method_dto.CreateForumTopic[ID]) (*forum_dto.ForumTopic, error) {
	data, err := json.Marshal(cft)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/createForumTopic", data)
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

// https://core.telegram.org/bots/api#editforumtopic
func EditForumTopic[ID int64 | string](ctx context.Context, eft method_dto.EditForumTopic[ID]) (*bool, error) {
	data, err := json.Marshal(eft)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editForumTopic", data)
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

// https://core.telegram.org/bots/api#closeforumtopic
func CloseForumTopic[ID int64 | string](ctx context.Context, cft method_dto.CloseForumTopic[ID]) (*bool, error) {
	data, err := json.Marshal(cft)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/closeForumTopic", data)
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

// https://core.telegram.org/bots/api#reopenforumtopic
func ReopenForumTopic[ID int64 | string](ctx context.Context, rft method_dto.ReopenForumTopic[ID]) (*bool, error) {
	data, err := json.Marshal(rft)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/reopenForumTopic", data)
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

// https://core.telegram.org/bots/api#deleteforumtopic
func DeleteForumTopic[ID int64 | string](ctx context.Context, dft method_dto.DeleteForumTopic[ID]) (*bool, error) {
	data, err := json.Marshal(dft)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteForumTopic", data)
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

// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func UnpinAllForumTopicMessages[ID int64 | string](ctx context.Context, uaftm method_dto.UnpinAllForumTopicMessages[ID]) (*bool, error) {
	data, err := json.Marshal(uaftm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinAllForumTopicMessages", data)
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

// https://core.telegram.org/bots/api#editgeneralforumtopic
func EditGeneralForumTopic[ID int64 | string](ctx context.Context, egft method_dto.EditGeneralForumTopic[ID]) (*bool, error) {
	data, err := json.Marshal(egft)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editGeneralForumTopic", data)
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

// https://core.telegram.org/bots/api#closegeneralforumtopic
func CloseGeneralForumTopic[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/closeGeneralForumTopic", data)
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

// https://core.telegram.org/bots/api#reopengeneralforumtopic
func ReopenGeneralForumTopic[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/reopenGeneralForumTopic", data)
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

// https://core.telegram.org/bots/api#hidegeneralforumtopic
func HideGeneralForumTopic[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/hideGeneralForumTopic", data)
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

// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func UnhideGeneralForumTopic[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unhideGeneralForumTopic", data)
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

// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func UnpinAllGeneralForumTopicMessages[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinAllGeneralForumTopicMessages", data)
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
