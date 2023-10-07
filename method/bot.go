package method

import (
	"context"
	"encoding/json"
	"log/slog"

	bot_dto "github.com/Avirtan/TGoBot/dto/bot"
	message_dto "github.com/Avirtan/TGoBot/dto/message"
	method_dto "github.com/Avirtan/TGoBot/dto/method"
	user_dto "github.com/Avirtan/TGoBot/dto/user"
	utils_dto "github.com/Avirtan/TGoBot/dto/utils"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#getme
func GetMe(ctx context.Context) (*user_dto.User, error) {
	response, err := request.RequestWithContext(ctx, request.Get, GetUrl()+"/GetMe")
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[user_dto.User](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func GetMyDefaultAdministratorRights(ctx context.Context, data method_dto.GetMyDefaultAdministratorRights) (*message_dto.ChatAdministratorRights, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyDefaultAdministratorRights", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.ChatAdministratorRights](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func SetMyDefaultAdministratorRights(ctx context.Context, data method_dto.SetMyDefaultAdministratorRights) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyDefaultAdministratorRights", marshalBytes)
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

// https://core.telegram.org/bots/api#getchatmenubutton
func GetChatMenuButton(ctx context.Context, data method_dto.GetChatMenuButton) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatMenuButton", marshalBytes)
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

// https://core.telegram.org/bots/api#setchatmenubutton
func SetChatMenuButton(ctx context.Context, data method_dto.SetChatMenuButton) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatMenuButton", marshalBytes)
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

// https://core.telegram.org/bots/api#getmyshortdescription
func GetMyShortDescription(ctx context.Context, data method_dto.GetMyShortDescription) (*bot_dto.BotShortDescription, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyShortDescription", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bot_dto.BotShortDescription](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#setmyshortdescription
func SetMyShortDescription(ctx context.Context, data method_dto.SetMyShortDescription) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyShortDescription", marshalBytes)
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

// https://core.telegram.org/bots/api#getmydescription
func GetMyDescription(ctx context.Context, data method_dto.GetMyDescription) (*bot_dto.BotDescription, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyDescription", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bot_dto.BotDescription](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#setmydescription
func SetMyDescription(ctx context.Context, data method_dto.SetMyDescription) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyDescription", marshalBytes)
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

// https://core.telegram.org/bots/api#getmyname
func GetMyName(ctx context.Context, data method_dto.GetMyName) (*bot_dto.BotName, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyName", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[bot_dto.BotName](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#setmyname
func SetMyName(ctx context.Context, data method_dto.SetMyName) error {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyName", marshalBytes)
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

// https://core.telegram.org/bots/api#getfile
func GetFile(ctx context.Context, data method_dto.GetFile) (*utils_dto.File, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getFile", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[utils_dto.File](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}
