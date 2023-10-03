package method

import (
	"TGoBot/dto"
	game_dto "TGoBot/dto/game"
	message_dto "TGoBot/dto/message"
	method_dto "TGoBot/dto/method"
	"TGoBot/request"
	"context"
	"encoding/json"
	"log/slog"
)

// https://core.telegram.org/bots/api#sendgame
func SendGame(ctx context.Context, sg method_dto.SendGame) (*message_dto.Message, error) {
	data, err := json.Marshal(sg)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendGame", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[message_dto.Message](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#setgamescore
func SetGameScore(ctx context.Context, sgs method_dto.SetGameScore) (*dto.Response, error) {
	data, err := json.Marshal(sgs)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setGameScore", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandler(response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#getgamehighscores
func GetGameHighScores(ctx context.Context, gghs method_dto.GetGameHighScores) (*[]game_dto.GameHighScore, error) {
	data, err := json.Marshal(gghs)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getGameHighScores", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[[]game_dto.GameHighScore](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}
