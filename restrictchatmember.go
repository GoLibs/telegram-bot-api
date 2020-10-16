package go_telegram_bot_api

import (
	"encoding/json"
	"time"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type restrictChatMember struct {
	parent      *TelegramBot
	chatId      interface{}
	userId      int64
	permissions structs.ChatPermissions
	untilDate   int64
}

func (rcm *restrictChatMember) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId      interface{}             `json:"chat_id"`
		UserId      int64                   `json:"user_id"`
		Permissions structs.ChatPermissions `json:"permissions"`
		UntilDate   int64                   `json:"until_date,omitempty"`
	}{
		ChatId:      rcm.chatId,
		UserId:      rcm.userId,
		Permissions: rcm.permissions,
		UntilDate:   rcm.untilDate,
	})
}

func (rcm *restrictChatMember) response() interface{} {
	var resp bool
	return &resp
}

func (rcm *restrictChatMember) method() string {
	return "POST"
}

func (rcm *restrictChatMember) endpoint() string {
	return "restrictChatMember"
}

func (rcm *restrictChatMember) SetUserId(userId int64) *restrictChatMember {
	rcm.userId = userId
	return rcm
}

func (rcm *restrictChatMember) SetChatId(chatId int64) *restrictChatMember {
	rcm.chatId = chatId
	return rcm
}

func (rcm *restrictChatMember) SetChatUsername(username string) *restrictChatMember {
	rcm.chatId = username
	return rcm
}

func (rcm *restrictChatMember) SetUntilDate(untilDate int64) *restrictChatMember {
	rcm.untilDate = untilDate
	return rcm
}

func (rcm *restrictChatMember) SetPermanent() *restrictChatMember {
	rcm.untilDate = time.Now().Add(-(time.Second * 31)).Unix()
	return rcm
}

func (rcm *restrictChatMember) DisallowMessages() *restrictChatMember {
	rcm.permissions.CanSendMessages = false
	return rcm
}

func (rcm *restrictChatMember) DisallowMediaMessages() *restrictChatMember {
	rcm.permissions.CanSendMediaMessages = false
	return rcm
}

func (rcm *restrictChatMember) DisallowPolls() *restrictChatMember {
	rcm.permissions.CanSendPolls = false
	return rcm
}

func (rcm *restrictChatMember) DisallowOtherMessages() *restrictChatMember {
	rcm.permissions.CanSendOtherMessages = false
	return rcm
}

func (rcm *restrictChatMember) DisallowWebPagePreviews() *restrictChatMember {
	rcm.permissions.CanAddWebPagePreviews = false
	return rcm
}

func (rcm *restrictChatMember) DisallowChangeInfo() *restrictChatMember {
	rcm.permissions.CanChangeInfo = false
	return rcm
}

func (rcm *restrictChatMember) DisallowInviteUsers() *restrictChatMember {
	rcm.permissions.CanInviteUsers = false
	return rcm
}

func (rcm *restrictChatMember) DisallowPinMessages() *restrictChatMember {
	rcm.permissions.CanPinMessages = false
	return rcm
}

func (rcm *restrictChatMember) AllowMessages() *restrictChatMember {
	rcm.permissions.CanSendMessages = true
	return rcm
}

func (rcm *restrictChatMember) AllowMediaMessages() *restrictChatMember {
	rcm.permissions.CanSendMediaMessages = true
	return rcm
}

func (rcm *restrictChatMember) AllowPolls() *restrictChatMember {
	rcm.permissions.CanSendPolls = true
	return rcm
}

func (rcm *restrictChatMember) AllowOtherMessages() *restrictChatMember {
	rcm.permissions.CanSendOtherMessages = true
	return rcm
}

func (rcm *restrictChatMember) AllowWebPagePreviews() *restrictChatMember {
	rcm.permissions.CanAddWebPagePreviews = true
	return rcm
}

func (rcm *restrictChatMember) AllowChangeInfo() *restrictChatMember {
	rcm.permissions.CanChangeInfo = true
	return rcm
}

func (rcm *restrictChatMember) AllowInviteUsers() *restrictChatMember {
	rcm.permissions.CanInviteUsers = true
	return rcm
}

func (rcm *restrictChatMember) AllowPinMessages() *restrictChatMember {
	rcm.permissions.CanPinMessages = true
	return rcm
}

/*func (rcm *restrictChatMember) Send() (message *structs.Message, err error) {
	message, err = rcm.parent.Send(sd)
	return
}
*/
