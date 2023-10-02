package method

import (
	method_dto "TGoBot/dto/method"
	"TGoBot/request"
	"context"
	"encoding/json"
	"log/slog"
)

// https://core.telegram.org/bots/api#sendphoto
func SendPhoto[ID int64 | string](ctx context.Context, msg method_dto.SendPhoto[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendPhoto", data)
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

// https://core.telegram.org/bots/api#sendaudio
func SendAudio[ID int64 | string](ctx context.Context, msg method_dto.SendAudio[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendAudio", data)
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

// https://core.telegram.org/bots/api#senddocument
func SendDocument[ID int64 | string](ctx context.Context, msg method_dto.SendDocument[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendDocument", data)
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

// https://core.telegram.org/bots/api#sendvideo
func SendVideo[ID int64 | string](ctx context.Context, msg method_dto.SendVideo[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendVideo", data)
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

// https://core.telegram.org/bots/api#sendanimation
func SendAnimation[ID int64 | string](ctx context.Context, msg method_dto.SendAnimation[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendAnimation", data)
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

// https://core.telegram.org/bots/api#sendvoice
func SendVoice[ID int64 | string](ctx context.Context, msg method_dto.SendVoice[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendVoice", data)
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

// https://core.telegram.org/bots/api#sendvideonote
func SendVideoNote[ID int64 | string](ctx context.Context, msg method_dto.SendVideoNote[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendVideoNote", data)
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

// https://core.telegram.org/bots/api#sendmediagroup
func SendMediaGroup[ID int64 | string](ctx context.Context, msg method_dto.SendMediaGroup[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendMediaGroup", data)
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

// https://core.telegram.org/bots/api#sendlocation
func SendLocation[ID int64 | string](ctx context.Context, msg method_dto.SendLocation[ID]) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendLocation", data)
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
