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
	responseJson, err := request.ResponseHandlerToType[user_dto.User](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func GetMyDefaultAdministratorRights(ctx context.Context, r method_dto.GetMyDefaultAdministratorRights) (*message_dto.ChatAdministratorRights, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyDefaultAdministratorRights", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[message_dto.ChatAdministratorRights](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func SetMyDefaultAdministratorRights(ctx context.Context, r method_dto.SetMyDefaultAdministratorRights) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyDefaultAdministratorRights", data)
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

// https://core.telegram.org/bots/api#getchatmenubutton
func GetChatMenuButton(ctx context.Context, chat method_dto.GetChatMenuButton) error {
	data, err := json.Marshal(chat)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatMenuButton", data)
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

// https://core.telegram.org/bots/api#setchatmenubutton
func SetChatMenuButton(ctx context.Context, chat method_dto.SetChatMenuButton) error {
	data, err := json.Marshal(chat)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatMenuButton", data)
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

// https://core.telegram.org/bots/api#getmyshortdescription
func GetMyShortDescription(ctx context.Context, desc method_dto.GetMyShortDescription) (*bot_dto.BotShortDescription, error) {
	data, err := json.Marshal(desc)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyShortDescription", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[bot_dto.BotShortDescription](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

// https://core.telegram.org/bots/api#setmyshortdescription
func SetMyShortDescription(ctx context.Context, desc method_dto.SetMyShortDescription) error {
	data, err := json.Marshal(desc)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyShortDescription", data)
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

// https://core.telegram.org/bots/api#getmydescription
func GetMyDescription(ctx context.Context, desc method_dto.GetMyDescription) (*bot_dto.BotDescription, error) {
	data, err := json.Marshal(desc)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyDescription", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[bot_dto.BotDescription](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

// https://core.telegram.org/bots/api#setmydescription
func SetMyDescription(ctx context.Context, desc method_dto.SetMyDescription) error {
	data, err := json.Marshal(desc)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyDescription", data)
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

// https://core.telegram.org/bots/api#getmyname
func GetMyName(ctx context.Context, desc method_dto.GetMyName) (*bot_dto.BotName, error) {
	data, err := json.Marshal(desc)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getMyName", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[bot_dto.BotName](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}

// https://core.telegram.org/bots/api#setmyname
func SetMyName(ctx context.Context, desc method_dto.SetMyName) error {
	data, err := json.Marshal(desc)
	if err != nil {
		return err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setMyName", data)
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

// https://core.telegram.org/bots/api#getfile
func GetFile(ctx context.Context, desc method_dto.GetFile) (*utils_dto.File, error) {
	data, err := json.Marshal(desc)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getFile", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[utils_dto.File](response)
	if err != nil {
		return nil, err
	}
	slog.Debug("info", "response", responseJson)
	return responseJson, nil
}
