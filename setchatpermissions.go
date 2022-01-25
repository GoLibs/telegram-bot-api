package tgbotapi

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type setChatPermissions struct {
	parent      *TelegramBot
	chatId      interface{}
	permissions structs.ChatPermissions
	customTitle string
}

func (scp *setChatPermissions) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId      interface{}             `json:"chat_id"`
		UserId      int64                   `json:"user_id"`
		Permissions structs.ChatPermissions `json:"permissions"`
	}{
		ChatId:      scp.chatId,
		Permissions: scp.permissions,
	})
}

func (scp *setChatPermissions) response() interface{} {
	return structs.ResponseTypeBool()
}

func (scp *setChatPermissions) method() string {
	return "POST"
}

func (scp *setChatPermissions) endpoint() string {
	return "setChatPermissions"
}

func (scp *setChatPermissions) SetChatId(chatId int64) *setChatPermissions {
	scp.chatId = chatId
	return scp
}

func (scp *setChatPermissions) SetChatUsername(username string) *setChatPermissions {
	scp.chatId = username
	return scp
}

func (scp *setChatPermissions) EnableChangeInfo() *setChatPermissions {
	scp.permissions.CanChangeInfo = true
	return scp
}

func (scp *setChatPermissions) EnableSendMessages() *setChatPermissions {
	scp.permissions.CanSendMessages = true
	return scp
}

func (scp *setChatPermissions) EnableSendMediaMessages() *setChatPermissions {
	scp.permissions.CanSendMediaMessages = true
	return scp
}

func (scp *setChatPermissions) EnableSendPolls() *setChatPermissions {
	scp.permissions.CanSendPolls = true
	return scp
}

func (scp *setChatPermissions) EnableSendOtherMessages() *setChatPermissions {
	scp.permissions.CanSendOtherMessages = true
	return scp
}

func (scp *setChatPermissions) EnableWebPagePreviews() *setChatPermissions {
	scp.permissions.CanAddWebPagePreviews = true
	return scp
}

func (scp *setChatPermissions) EnableInviteUsers() *setChatPermissions {
	scp.permissions.CanInviteUsers = true
	return scp
}

func (scp *setChatPermissions) EnablePinMessages() *setChatPermissions {
	scp.permissions.CanChangeInfo = true
	return scp
}

func (scp *setChatPermissions) DisableChangeInfo() *setChatPermissions {
	scp.permissions.CanChangeInfo = false
	return scp
}

func (scp *setChatPermissions) DisableSendMessages() *setChatPermissions {
	scp.permissions.CanSendMessages = false
	return scp
}

func (scp *setChatPermissions) DisableSendMediaMessages() *setChatPermissions {
	scp.permissions.CanSendMediaMessages = false
	return scp
}

func (scp *setChatPermissions) DisableSendPolls() *setChatPermissions {
	scp.permissions.CanSendPolls = false
	return scp
}

func (scp *setChatPermissions) DisableSendOtherMessages() *setChatPermissions {
	scp.permissions.CanSendOtherMessages = false
	return scp
}

func (scp *setChatPermissions) DisableWebPagePreviews() *setChatPermissions {
	scp.permissions.CanAddWebPagePreviews = false
	return scp
}

func (scp *setChatPermissions) DisableInviteUsers() *setChatPermissions {
	scp.permissions.CanInviteUsers = false
	return scp
}

func (scp *setChatPermissions) DisablePinMessages() *setChatPermissions {
	scp.permissions.CanChangeInfo = false
	return scp
}
