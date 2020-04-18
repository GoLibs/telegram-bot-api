package go_telegram_bot_api

import (
	"encoding/json"
)

type promoteChatMember struct {
	parent             *TelegramBot
	chatId             interface{}
	userId             int64
	canChangeInfo      bool
	canPostMessages    bool
	canEditMessages    bool
	canDeleteMessages  bool
	canInviteUsers     bool
	canRestrictMembers bool
	canPinMessages     bool
	canPromoteMembers  bool
}

func (pcm *promoteChatMember) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId             interface{} `json:"chat_id"`
		UserId             int64       `json:"user_id"`
		CanChangeInfo      bool        `json:"can_change_info"`
		CanPostMessages    bool        `json:"can_post_messages"`
		CanEditMessages    bool        `json:"can_edit_messages"`
		CanDeleteMessages  bool        `json:"can_delete_messages"`
		CanInviteUsers     bool        `json:"can_invite_users"`
		CanRestrictMembers bool        `json:"can_restrict_members"`
		CanPinMessages     bool        `json:"can_pin_messages"`
		CanPromoteMembers  bool        `json:"can_promote_members"`
	}{
		ChatId:             pcm.chatId,
		UserId:             pcm.userId,
		CanChangeInfo:      pcm.canChangeInfo,
		CanPostMessages:    pcm.canPostMessages,
		CanEditMessages:    pcm.canEditMessages,
		CanDeleteMessages:  pcm.canDeleteMessages,
		CanInviteUsers:     pcm.canInviteUsers,
		CanRestrictMembers: pcm.canRestrictMembers,
		CanPinMessages:     pcm.canPinMessages,
		CanPromoteMembers:  pcm.canPromoteMembers,
	})
}

func (pcm *promoteChatMember) response() interface{} {
	var resp bool
	return &resp
}

func (pcm *promoteChatMember) method() string {
	return "POST"
}

func (pcm *promoteChatMember) endpoint() string {
	return "promoteChatMember"
}

func (pcm *promoteChatMember) SetUserId(userId int64) *promoteChatMember {
	pcm.userId = userId
	return pcm
}

func (pcm *promoteChatMember) SetChatId(chatId int64) *promoteChatMember {
	pcm.chatId = chatId
	return pcm
}

func (pcm *promoteChatMember) SetChatUsername(username string) *promoteChatMember {
	pcm.chatId = username
	return pcm
}

func (pcm *promoteChatMember) EnableChangeInfo() *promoteChatMember {
	pcm.canChangeInfo = true
	return pcm
}

func (pcm *promoteChatMember) EnablePostMessages() *promoteChatMember {
	pcm.canPostMessages = true
	return pcm
}

func (pcm *promoteChatMember) EnableEditMessages() *promoteChatMember {
	pcm.canEditMessages = true
	return pcm
}

func (pcm *promoteChatMember) EnableDeleteMessages() *promoteChatMember {
	pcm.canDeleteMessages = true
	return pcm
}

func (pcm *promoteChatMember) EnableInviteUsers() *promoteChatMember {
	pcm.canInviteUsers = true
	return pcm
}

func (pcm *promoteChatMember) EnableRestrictMembers() *promoteChatMember {
	pcm.canRestrictMembers = true
	return pcm
}

func (pcm *promoteChatMember) EnablePinMessages() *promoteChatMember {
	pcm.canChangeInfo = true
	return pcm
}

func (pcm *promoteChatMember) EnablePromoteMembers() *promoteChatMember {
	pcm.canPromoteMembers = true
	return pcm
}

func (pcm *promoteChatMember) DisableChangeInfo() *promoteChatMember {
	pcm.canChangeInfo = false
	return pcm
}

func (pcm *promoteChatMember) DisablePostMessages() *promoteChatMember {
	pcm.canPostMessages = false
	return pcm
}

func (pcm *promoteChatMember) DisableEditMessages() *promoteChatMember {
	pcm.canEditMessages = false
	return pcm
}

func (pcm *promoteChatMember) DisableDeleteMessages() *promoteChatMember {
	pcm.canDeleteMessages = false
	return pcm
}

func (pcm *promoteChatMember) DisableInviteUsers() *promoteChatMember {
	pcm.canInviteUsers = false
	return pcm
}

func (pcm *promoteChatMember) DisableRestrictMembers() *promoteChatMember {
	pcm.canRestrictMembers = false
	return pcm
}

func (pcm *promoteChatMember) DisablePinMessages() *promoteChatMember {
	pcm.canChangeInfo = false
	return pcm
}

func (pcm *promoteChatMember) DisablePromoteMembers() *promoteChatMember {
	pcm.canPromoteMembers = false
	return pcm
}
