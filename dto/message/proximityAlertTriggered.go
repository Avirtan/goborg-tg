package message_dto

import user_dto "github.com/Avirtan/goborg-tg/dto/user"

// https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler user_dto.User `json:"traveler"`
	Watcher  user_dto.User `json:"watcher"`
	Distance int64         `json:"distance"`
}
