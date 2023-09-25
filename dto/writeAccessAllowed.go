package dto

// https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
}
