package message_dto

import (
	message_entity_dto "github.com/Avirtan/TGoBot/dto/message_entity"
	user_dto "github.com/Avirtan/TGoBot/dto/user"
)

// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollId    string         `json:"poll_id"`
	VoterChat *Chat          `json:"voter_chat,omitempty"`
	User      *user_dto.User `json:"user,omitempty"`
	OptionIds []int          `json:"option_ids"`
}

// https://core.telegram.org/bots/api#poll
type Poll struct {
	Id                    string                              `json:"id"`
	Question              string                              `json:"question"`
	Explanation           string                              `json:"explanation,omitempty"`
	Type                  string                              `json:"type"`
	TotalVoterCount       int                                 `json:"total_voter_count"`
	CorrectOptionId       int                                 `json:"correct_option_id,omitempty"`
	OpenPeriod            int                                 `json:"open_period,omitempty"`
	CloseDate             int                                 `json:"close_date,omitempty"`
	ExplanationEntities   []*message_entity_dto.MessageEntity `json:"explanation_entities,omitempty"`
	Options               []PollOption                        `json:"options"`
	IsClosed              bool                                `json:"is_closed"`
	IsAnonymous           bool                                `json:"is_anonymous"`
	AllowsMultipleAnswers bool                                `json:"allows_multiple_answers"`
}
