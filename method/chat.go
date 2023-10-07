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
func BanChatMember[ID int64 | string](ctx context.Context, bcm method_dto.BanChatMember[ID]) (*bool, error) {
	data, err := json.Marshal(bcm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/banChatMember", data)
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

// https://core.telegram.org/bots/api#unbanchatmember
func UnbanChatMember[ID int64 | string](ctx context.Context, ucm method_dto.UnbanChatMember[ID]) (*bool, error) {
	data, err := json.Marshal(ucm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unbanChatMember", data)
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

// https://core.telegram.org/bots/api#restrictchatmember
func RestrictChatMember[ID int64 | string](ctx context.Context, rcm method_dto.RestrictChatMember[ID]) (*bool, error) {
	data, err := json.Marshal(rcm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/restrictChatMember", data)
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

// https://core.telegram.org/bots/api#promotechatmember
func PromoteChatMember[ID int64 | string](ctx context.Context, pcm method_dto.PromoteChatMember[ID]) (*bool, error) {
	data, err := json.Marshal(pcm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/promoteChatMember", data)
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

// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func SetChatAdministratorCustomTitle[ID int64 | string](ctx context.Context, s method_dto.SetChatAdministratorCustomTitle[ID]) (*bool, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatAdministratorCustomTitle", data)
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

// https://core.telegram.org/bots/api#banchatsenderchat
func BanChatSenderChat[ID int64 | string](ctx context.Context, bcsc method_dto.BanChatSenderChat[ID]) (*bool, error) {
	data, err := json.Marshal(bcsc)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/banChatSenderChat", data)
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

// https://core.telegram.org/bots/api#unbanchatsenderchat
func UnbanChatSenderChat[ID int64 | string](ctx context.Context, ucsc method_dto.UnbanChatSenderChat[ID]) (*bool, error) {
	data, err := json.Marshal(ucsc)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unbanChatSenderChat", data)
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

// https://core.telegram.org/bots/api#setchatpermissions
func SetChatPermissions[ID int64 | string](ctx context.Context, scp method_dto.SetChatPermissions[ID]) (*bool, error) {
	data, err := json.Marshal(scp)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatPermissions", data)
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

// https://core.telegram.org/bots/api#exportchatinvitelink
func ExportChatInviteLink[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*string, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/exportChatInviteLink", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[string](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#createchatinvitelink
func CreateChatInviteLink[ID int64 | string](ctx context.Context, ccil method_dto.CreateChatInviteLink[ID]) (*message_dto.ChatInviteLink, error) {
	data, err := json.Marshal(ccil)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/createChatInviteLink", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[message_dto.ChatInviteLink](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#editchatinvitelink
func EditChatInviteLink[ID int64 | string](ctx context.Context, ecil method_dto.EditChatInviteLink[ID]) (*message_dto.ChatInviteLink, error) {
	data, err := json.Marshal(ecil)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/editChatInviteLink", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[message_dto.ChatInviteLink](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#revokechatinvitelink
func RevokeChatInviteLink[ID int64 | string](ctx context.Context, rcil method_dto.RevokeChatInviteLink[ID]) (*message_dto.ChatInviteLink, error) {
	data, err := json.Marshal(rcil)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/revokeChatInviteLink", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[message_dto.ChatInviteLink](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
func ApproveChatJoinRequest[ID int64 | string](ctx context.Context, acjr method_dto.ApproveChatJoinRequest[ID]) (*bool, error) {
	data, err := json.Marshal(acjr)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/approveChatJoinRequest", data)
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

// https://core.telegram.org/bots/api#declinechatjoinrequest
func DeclineChatJoinRequest[ID int64 | string](ctx context.Context, acjr method_dto.DeclineChatJoinRequest[ID]) (*bool, error) {
	data, err := json.Marshal(acjr)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/declineChatJoinRequest", data)
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

// https://core.telegram.org/bots/api#setchatphoto
func SetChatPhoto[ID int64 | string](ctx context.Context, scp method_dto.SetChatPhoto[ID]) (*bool, error) {
	data, err := json.Marshal(scp)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatPhoto", data)
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

// https://core.telegram.org/bots/api#deletechatphoto
func DeleteChatPhoto[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteChatPhoto", data)
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

// https://core.telegram.org/bots/api#setchattitle
func SetChatTitle[ID int64 | string](ctx context.Context, sct method_dto.SetChatTitle[ID]) (*bool, error) {
	data, err := json.Marshal(sct)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatTitle", data)
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

// https://core.telegram.org/bots/api#setchatdescription
func SetChatDescription[ID int64 | string](ctx context.Context, scd method_dto.SetChatDescription[ID]) (*bool, error) {
	data, err := json.Marshal(scd)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatDescription", data)
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

// https://core.telegram.org/bots/api#pinchatmessage
func PinChatMessage[ID int64 | string](ctx context.Context, pcm method_dto.PinChatMessage[ID]) (*bool, error) {
	data, err := json.Marshal(pcm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/pinChatMessage", data)
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

// https://core.telegram.org/bots/api#unpinchatmessage
func UnpinChatMessage[ID int64 | string](ctx context.Context, pcm method_dto.UnpinChatMessage[ID]) (*bool, error) {
	data, err := json.Marshal(pcm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinChatMessage", data)
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

// https://core.telegram.org/bots/api#unpinallchatmessages
func UnpinAllChatMessages[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/unpinAllChatMessages", data)
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

// https://core.telegram.org/bots/api#leavechat
func LeaveChat[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/leaveChat", data)
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

// https://core.telegram.org/bots/api#getchat
func GetChat[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*message_dto.Chat, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChat", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[message_dto.Chat](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#getchatadministrators
func GetChatAdministrators[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*[]message_dto.ChatMember, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatAdministrators", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[[]message_dto.ChatMember](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#getchatmembercount
func GetChatMemberCount[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*int64, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatMemberCount", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[int64](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#getchatmember
func GetChatMember[ID int64 | string](ctx context.Context, gcm method_dto.GetChatMember[ID]) (*message_dto.ChatMember, error) {
	data, err := json.Marshal(gcm)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/getChatMember", data)
	if err != nil {
		return nil, err
	}

	responseJson, err := request.ResponseHandlerToType[message_dto.ChatMember](response)
	if err != nil {
		return nil, err
	}

	slog.Debug("info", "response", responseJson)

	return responseJson, nil
}

// https://core.telegram.org/bots/api#setchatstickerset
func SetChatStickerSet[ID int64 | string](ctx context.Context, scss method_dto.SetChatStickerSet[ID]) (*bool, error) {
	data, err := json.Marshal(scss)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/setChatStickerSet", data)
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

// https://core.telegram.org/bots/api#deletechatstickerset
func DeleteChatStickerSet[ID int64 | string](ctx context.Context, chatID method_dto.ChatID[ID]) (*bool, error) {
	data, err := json.Marshal(chatID)
	if err != nil {
		return nil, err
	}

	response, err := request.RequestWithContextAndData(ctx, request.Get, GetUrl()+"/deleteChatStickerSet", data)
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
