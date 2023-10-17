package method

import (
	"context"
	"encoding/json"
	"log/slog"

	message_dto "github.com/Avirtan/goborg-tg/dto/message"
	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	stickers_dto "github.com/Avirtan/goborg-tg/dto/stickers"
	utils_dto "github.com/Avirtan/goborg-tg/dto/utils"
	"github.com/Avirtan/goborg-tg/request"
)

// https://core.telegram.org/bots/api#sendsticker
func SendSticker[ID int64 | string](ctx context.Context, data method_dto.SendSticker[ID]) (*message_dto.Message, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/sendSticker", marshalBytes)
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

// https://core.telegram.org/bots/api#getstickerset
func GetStickerSet(ctx context.Context, data method_dto.GetStickerSet) (*stickers_dto.StickerSet, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getStickerSet", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[stickers_dto.StickerSet](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#getcustomemojistickers
func GetCustomEmojiStickers(ctx context.Context, data method_dto.GetCustomEmojiStickers) (*[]stickers_dto.Sticker, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getCustomEmojiStickers", marshalBytes)
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

// https://core.telegram.org/bots/api#uploadstickerfile
func UploadStickerFile(ctx context.Context, data method_dto.UploadStickerFile) (*utils_dto.File, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/uploadStickerFile", marshalBytes)
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

// https://core.telegram.org/bots/api#createnewstickerset
func CreateNewStickerSet(ctx context.Context, data method_dto.CreateNewStickerSet) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/createNewStickerSet", marshalBytes)
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

// https://core.telegram.org/bots/api#addstickertoset
func AddStickerToSet(ctx context.Context, data method_dto.AddStickerToSet) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/addStickerToSet", marshalBytes)
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

// https://core.telegram.org/bots/api#setstickerpositioninset
func SetStickerPositionInSet(ctx context.Context, data method_dto.SetStickerPositionInSet) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerPositionInSet", marshalBytes)
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

// https://core.telegram.org/bots/api#deletestickerfromset
func DeleteStickerFromSet(ctx context.Context, data method_dto.DeleteStickerFromSet) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteStickerFromSet", marshalBytes)
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

// https://core.telegram.org/bots/api#setstickeremojilist
func SetStickerEmojiList(ctx context.Context, data method_dto.SetStickerEmojiList) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerEmojiList", marshalBytes)
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

// https://core.telegram.org/bots/api#setstickerkeywords
func SetStickerKeywords(ctx context.Context, data method_dto.SetStickerKeywords) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerKeywords", marshalBytes)
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

// https://core.telegram.org/bots/api#setstickermaskposition
func SetStickerMaskPosition(ctx context.Context, data method_dto.SetStickerMaskPosition) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerMaskPosition", marshalBytes)
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

// https://core.telegram.org/bots/api#setstickersettitle
func SetStickerSetTitle(ctx context.Context, data method_dto.SetStickerSetTitle) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerSetTitle", marshalBytes)
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

// https://core.telegram.org/bots/api#setstickersetthumbnail
func SetStickerSetThumbnail(ctx context.Context, data method_dto.SetStickerSetThumbnail) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setStickerSetThumbnail", marshalBytes)
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

// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func SetCustomEmojiStickerSetThumbnail(ctx context.Context, data method_dto.SetCustomEmojiStickerSetThumbnail) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setCustomEmojiStickerSetThumbnail", marshalBytes)
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

// https://core.telegram.org/bots/api#deletestickerset
func DeleteStickerSet(ctx context.Context, data method_dto.DeleteStickerSet) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteStickerSet", marshalBytes)
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
