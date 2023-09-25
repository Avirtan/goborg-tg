package dto

// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}
