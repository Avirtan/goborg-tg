package method

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/Avirtan/TGoBot/dto"
	game_dto "github.com/Avirtan/TGoBot/dto/game"
	message_dto "github.com/Avirtan/TGoBot/dto/message"
	method_dto "github.com/Avirtan/TGoBot/dto/method"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#sendgame
func SendGame(ctx context.Context, data method_dto.SendGame) (*message_dto.Message, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendGame", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.Message](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#setgamescore
func SetGameScore(ctx context.Context, data method_dto.SetGameScore) (*dto.Response, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setGameScore", marshalBytes)
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

// https://core.telegram.org/bots/api#getgamehighscores
func GetGameHighScores(ctx context.Context, data method_dto.GetGameHighScores) (*[]game_dto.GameHighScore, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getGameHighScores", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[[]game_dto.GameHighScore](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}
