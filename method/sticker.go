package method

import (
	"context"
	"encoding/json"
	"log/slog"

	message_dto "github.com/Avirtan/TGoBot/dto/message"
	method_dto "github.com/Avirtan/TGoBot/dto/method"
	stickers_dto "github.com/Avirtan/TGoBot/dto/stickers"
	utils_dto "github.com/Avirtan/TGoBot/dto/utils"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#sendsticker
func SendSticker[ID int64 | string](ctx context.Context, ss method_dto.SendSticker[ID]) (*message_dto.Message, error) {
	data, err := json.Marshal(ss)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendSticker", data)
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

// https://core.telegram.org/bots/api#getstickerset
func GetStickerSet(ctx context.Context, gss method_dto.GetStickerSet) (*stickers_dto.StickerSet, error) {
	data, err := json.Marshal(gss)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getStickerSet", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[stickers_dto.StickerSet](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#getcustomemojistickers
func GetCustomEmojiStickers(ctx context.Context, gces method_dto.GetCustomEmojiStickers) (*[]stickers_dto.Sticker, error) {
	data, err := json.Marshal(gces)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getCustomEmojiStickers", data)
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

// https://core.telegram.org/bots/api#uploadstickerfile
func UploadStickerFile(ctx context.Context, usf method_dto.UploadStickerFile) (*utils_dto.File, error) {
	data, err := json.Marshal(usf)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/uploadStickerFile", data)
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

// https://core.telegram.org/bots/api#createnewstickerset
func CreateNewStickerSet(ctx context.Context, usf method_dto.CreateNewStickerSet) (*bool, error) {
	data, err := json.Marshal(usf)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/createNewStickerSet", data)
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

// https://core.telegram.org/bots/api#addstickertoset
func AddStickerToSet(ctx context.Context, asts method_dto.AddStickerToSet) (*bool, error) {
	data, err := json.Marshal(asts)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/addStickerToSet", data)
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

// https://core.telegram.org/bots/api#setstickerpositioninset
func SetStickerPositionInSet(ctx context.Context, sspis method_dto.SetStickerPositionInSet) (*bool, error) {
	data, err := json.Marshal(sspis)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerPositionInSet", data)
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

// https://core.telegram.org/bots/api#deletestickerfromset
func DeleteStickerFromSet(ctx context.Context, dsfs method_dto.DeleteStickerFromSet) (*bool, error) {
	data, err := json.Marshal(dsfs)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteStickerFromSet", data)
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

// https://core.telegram.org/bots/api#setstickeremojilist
func SetStickerEmojiList(ctx context.Context, ssel method_dto.SetStickerEmojiList) (*bool, error) {
	data, err := json.Marshal(ssel)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerEmojiList", data)
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

// https://core.telegram.org/bots/api#setstickerkeywords
func SetStickerKeywords(ctx context.Context, ssk method_dto.SetStickerKeywords) (*bool, error) {
	data, err := json.Marshal(ssk)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerKeywords", data)
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

// https://core.telegram.org/bots/api#setstickermaskposition
func SetStickerMaskPosition(ctx context.Context, ssmp method_dto.SetStickerMaskPosition) (*bool, error) {
	data, err := json.Marshal(ssmp)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerMaskPosition", data)
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

// https://core.telegram.org/bots/api#setstickersettitle
func SetStickerSetTitle(ctx context.Context, ssst method_dto.SetStickerSetTitle) (*bool, error) {
	data, err := json.Marshal(ssst)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerSetTitle", data)
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

// https://core.telegram.org/bots/api#setstickersetthumbnail
func SetStickerSetThumbnail(ctx context.Context, ssst method_dto.SetStickerSetThumbnail) (*bool, error) {
	data, err := json.Marshal(ssst)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerSetThumbnail", data)
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

// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func SetCustomEmojiStickerSetThumbnail(ctx context.Context, scesst method_dto.SetCustomEmojiStickerSetThumbnail) (*bool, error) {
	data, err := json.Marshal(scesst)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setCustomEmojiStickerSetThumbnail", data)
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

// https://core.telegram.org/bots/api#deletestickerset
func DeleteStickerSet(ctx context.Context, dss method_dto.DeleteStickerSet) (*bool, error) {
	data, err := json.Marshal(dss)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteStickerSet", data)
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
