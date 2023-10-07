package method

import (
	"context"
	"encoding/json"
	"log/slog"

	message_dto "github.com/Avirtan/TGoBot/dto/message"
	method_dto "github.com/Avirtan/TGoBot/dto/method"
	"github.com/Avirtan/TGoBot/request"
)

// https://core.telegram.org/bots/api#banchatmember
func BanChatMember[ID int64 | string](ctx context.Context, data method_dto.BanChatMember[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/banChatMember", marshalBytes)
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

// https://core.telegram.org/bots/api#unbanchatmember
func UnbanChatMember[ID int64 | string](ctx context.Context, data method_dto.UnbanChatMember[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unbanChatMember", marshalBytes)
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

// https://core.telegram.org/bots/api#restrictchatmember
func RestrictChatMember[ID int64 | string](ctx context.Context, data method_dto.RestrictChatMember[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/restrictChatMember", marshalBytes)
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

// https://core.telegram.org/bots/api#promotechatmember
func PromoteChatMember[ID int64 | string](ctx context.Context, data method_dto.PromoteChatMember[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/promoteChatMember", marshalBytes)
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

// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func SetChatAdministratorCustomTitle[ID int64 | string](ctx context.Context, data method_dto.SetChatAdministratorCustomTitle[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatAdministratorCustomTitle", marshalBytes)
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

// https://core.telegram.org/bots/api#banchatsenderchat
func BanChatSenderChat[ID int64 | string](ctx context.Context, data method_dto.BanChatSenderChat[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/banChatSenderChat", marshalBytes)
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

// https://core.telegram.org/bots/api#unbanchatsenderchat
func UnbanChatSenderChat[ID int64 | string](ctx context.Context, data method_dto.UnbanChatSenderChat[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unbanChatSenderChat", marshalBytes)
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

// https://core.telegram.org/bots/api#setchatpermissions
func SetChatPermissions[ID int64 | string](ctx context.Context, data method_dto.SetChatPermissions[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatPermissions", marshalBytes)
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

// https://core.telegram.org/bots/api#exportchatinvitelink
func ExportChatInviteLink[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*string, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/exportChatInviteLink", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[string](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#createchatinvitelink
func CreateChatInviteLink[ID int64 | string](ctx context.Context, data method_dto.CreateChatInviteLink[ID]) (*message_dto.ChatInviteLink, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/createChatInviteLink", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.ChatInviteLink](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#editchatinvitelink
func EditChatInviteLink[ID int64 | string](ctx context.Context, data method_dto.EditChatInviteLink[ID]) (*message_dto.ChatInviteLink, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editChatInviteLink", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.ChatInviteLink](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#revokechatinvitelink
func RevokeChatInviteLink[ID int64 | string](ctx context.Context, data method_dto.RevokeChatInviteLink[ID]) (*message_dto.ChatInviteLink, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/revokeChatInviteLink", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.ChatInviteLink](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
func ApproveChatJoinRequest[ID int64 | string](ctx context.Context, data method_dto.ApproveChatJoinRequest[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/approveChatJoinRequest", marshalBytes)
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

// https://core.telegram.org/bots/api#declinechatjoinrequest
func DeclineChatJoinRequest[ID int64 | string](ctx context.Context, data method_dto.DeclineChatJoinRequest[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/declineChatJoinRequest", marshalBytes)
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

// https://core.telegram.org/bots/api#setchatphoto
func SetChatPhoto[ID int64 | string](ctx context.Context, data method_dto.SetChatPhoto[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatPhoto", marshalBytes)
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

// https://core.telegram.org/bots/api#deletechatphoto
func DeleteChatPhoto[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteChatPhoto", marshalBytes)
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

// https://core.telegram.org/bots/api#setchattitle
func SetChatTitle[ID int64 | string](ctx context.Context, data method_dto.SetChatTitle[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatTitle", marshalBytes)
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

// https://core.telegram.org/bots/api#setchatdescription
func SetChatDescription[ID int64 | string](ctx context.Context, data method_dto.SetChatDescription[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatDescription", marshalBytes)
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

// https://core.telegram.org/bots/api#pinchatmessage
func PinChatMessage[ID int64 | string](ctx context.Context, data method_dto.PinChatMessage[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/pinChatMessage", marshalBytes)
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

// https://core.telegram.org/bots/api#unpinchatmessage
func UnpinChatMessage[ID int64 | string](ctx context.Context, data method_dto.UnpinChatMessage[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinChatMessage", marshalBytes)
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

// https://core.telegram.org/bots/api#unpinallchatmessages
func UnpinAllChatMessages[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinAllChatMessages", marshalBytes)
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

// https://core.telegram.org/bots/api#leavechat
func LeaveChat[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/leaveChat", marshalBytes)
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

// https://core.telegram.org/bots/api#getchat
func GetChat[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*message_dto.Chat, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChat", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.Chat](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#getchatadministrators
func GetChatAdministrators[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*[]message_dto.ChatMember, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatAdministrators", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[[]message_dto.ChatMember](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#getchatmembercount
func GetChatMemberCount[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*int64, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatMemberCount", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[int64](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#getchatmember
func GetChatMember[ID int64 | string](ctx context.Context, data method_dto.GetChatMember[ID]) (*message_dto.ChatMember, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatMember", marshalBytes)
	if err != nil {
		return nil, err
	}

	responseData, err := request.ResponseHandlerToType[message_dto.ChatMember](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseData)

	return responseData, nil
}

// https://core.telegram.org/bots/api#setchatstickerset
func SetChatStickerSet[ID int64 | string](ctx context.Context, data method_dto.SetChatStickerSet[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatStickerSet", marshalBytes)
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

// https://core.telegram.org/bots/api#deletechatstickerset
func DeleteChatStickerSet[ID int64 | string](ctx context.Context, data method_dto.ChatID[ID]) (*bool, error) {
	marshalBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteChatStickerSet", marshalBytes)
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
