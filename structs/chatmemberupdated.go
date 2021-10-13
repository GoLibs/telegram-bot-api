package structs

type ChatMemberUpdated struct {
	Chat          *Chat           `json:"chat"`
	From          *User           `json:"from"`
	Date          int64           `json:"date"`
	OldChatMember *ChatMember     `json:"old_chat_member,omitempty"`
	NewChatMember *ChatMember     `json:"new_chat_member,omitempty"`
	InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}
