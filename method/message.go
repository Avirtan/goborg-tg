package method

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/Avirtan/goborg-tg/dto"
	message_dto "github.com/Avirtan/goborg-tg/dto/message"
	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	"github.com/Avirtan/goborg-tg/request"
)

// https://core.telegram.org/bots/api#sendmessage
func SendMessage[ID int64 | string](ctx context.Context, data method_dto.SendMessage[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendMessage", marshalBytes)
	if err != nil {
		return err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}

	slog.Debug("info", "response", responseData)

	return nil
}

// https://core.telegram.org/bots/api#forwardmessage
func ForwardMessage(ctx context.Context, data method_dto.ForwardMessage) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/forwardMessage", marshalBytes)
	if err != nil {
		return err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}

	slog.Debug("info", "response", responseData)

	return nil
}

// https://core.telegram.org/bots/api#copymessage
func CopyMessage(ctx context.Context, data method_dto.CopyMessage) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/copyMessage", marshalBytes)
	if err != nil {
		return err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}

	slog.Debug("info", "response", responseData)

	return nil
}

// https://core.telegram.org/bots/api#editmessagetext
func EditMessageText[ID int64 | string](ctx context.Context, data method_dto.EditMessageText[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editMessageText", marshalBytes)
	if err != nil {
		return err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return err
	}

	slog.Debug("info", "response", responseData)

	return nil
}

// https://core.telegram.org/bots/api#editmessagecaption
func EditMessageCaption[ID int64 | string](ctx context.Context, data method_dto.EditMessageCaption[ID]) (*dto.Response, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editMessageCaption", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#editmessagemedia
func EditMessageMedia[ID int64 | string](ctx context.Context, data method_dto.EditMessageMedia[ID]) (*dto.Response, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editMessageMedia", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil

}

// https://core.telegram.org/bots/api#editmessagelivelocation
func EditMessageLiveLocation[ID int64 | string](ctx context.Context, data method_dto.EditMessageLiveLocation[ID]) (*dto.Response, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editMessageLiveLocation", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#stopmessagelivelocation
func StopMessageLiveLocation[ID int64 | string](ctx context.Context, data method_dto.StopMessageLiveLocation[ID]) (*dto.Response, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/stopMessageLiveLocation", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
func EditMessageReplyMarkup[ID int64 | string](ctx context.Context, data method_dto.EditMessageReplyMarkup[ID]) (*dto.Response, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editMessageReplyMarkup", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandler(response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#stoppoll
func StopPoll[ID int64 | string](ctx context.Context, data method_dto.StopPoll[ID]) (*message_dto.Poll, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/stopPoll", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.Poll](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#deletemessage
func DeleteMessage[ID int64 | string](ctx context.Context, data method_dto.DeleteMessage[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteMessage", marshalBytes)
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
