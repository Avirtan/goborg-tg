package dto

// https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

// https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int64 `json:"duration"`
}

// https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}
