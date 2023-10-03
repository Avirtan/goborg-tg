package method_dto

import passport_dto "TGoBot/dto/passport"

// https://core.telegram.org/bots/api#setpassportdataerrors
type SetPassportDataErrors struct {
	UserID int                                 `json:"user_id"`
	Errors []passport_dto.PassportElementError `json:"errors"`
}

func NewSetPassportDataErrors(userID int, errors []passport_dto.PassportElementError) *SetPassportDataErrors {
	return &SetPassportDataErrors{
		UserID: userID,
		Errors: errors,
	}
}
