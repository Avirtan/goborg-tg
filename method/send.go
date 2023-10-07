package method

import (
	"context"
	"encoding/json"
	"log/slog"

	method_dto "github.com/Avirtan/TGoBot/dto/method"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#sendphoto
func SendPhoto[ID int64 | string](ctx context.Context, data method_dto.SendPhoto[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendPhoto", marshalBytes)
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

// https://core.telegram.org/bots/api#sendaudio
func SendAudio[ID int64 | string](ctx context.Context, data method_dto.SendAudio[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendAudio", marshalBytes)
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

// https://core.telegram.org/bots/api#senddocument
func SendDocument[ID int64 | string](ctx context.Context, data method_dto.SendDocument[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendDocument", marshalBytes)
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

// https://core.telegram.org/bots/api#sendvideo
func SendVideo[ID int64 | string](ctx context.Context, data method_dto.SendVideo[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendVideo", marshalBytes)
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

// https://core.telegram.org/bots/api#sendanimation
func SendAnimation[ID int64 | string](ctx context.Context, data method_dto.SendAnimation[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendAnimation", marshalBytes)
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

// https://core.telegram.org/bots/api#sendvoice
func SendVoice[ID int64 | string](ctx context.Context, data method_dto.SendVoice[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendVoice", marshalBytes)
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

// https://core.telegram.org/bots/api#sendvideonote
func SendVideoNote[ID int64 | string](ctx context.Context, data method_dto.SendVideoNote[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendVideoNote", marshalBytes)
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

// https://core.telegram.org/bots/api#sendmediagroup
func SendMediaGroup[ID int64 | string](ctx context.Context, data method_dto.SendMediaGroup[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendMediaGroup", marshalBytes)
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

// https://core.telegram.org/bots/api#sendlocation
func SendLocation[ID int64 | string](ctx context.Context, data method_dto.SendLocation[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendLocation", marshalBytes)
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

// https://core.telegram.org/bots/api#sendchataction
func SendChatAction[ID int64 | string](ctx context.Context, data method_dto.SendChatAction[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendChatAction", marshalBytes)
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

// https://core.telegram.org/bots/api#senddice
func SendDice[ID int64 | string](ctx context.Context, data method_dto.SendDice[ID]) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendDice", marshalBytes)
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
