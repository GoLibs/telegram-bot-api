package go_telegram_bot_api

import (
	"encoding/json"
)

type unbanChatMember struct {
	parent *TelegramBot
	chatId interface{}
	userId int64
}

func (ucm *unbanChatMember) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
		UserId int64       `json:"user_id"`
	}{
		ChatId: ucm.chatId,
		UserId: ucm.userId,
	})
}

func (ucm *unbanChatMember) response() interface{} {
	var resp bool
	return &resp
}

func (ucm *unbanChatMember) method() string {
	return "POST"
}

func (ucm *unbanChatMember) endpoint() string {
	return "unbanChatMember"
}

func (ucm *unbanChatMember) SetUserId(userId int64) *unbanChatMember {
	ucm.userId = userId
	return ucm
}

func (ucm *unbanChatMember) SetChatId(chatId int64) *unbanChatMember {
	ucm.chatId = chatId
	return ucm
}

func (ucm *unbanChatMember) SetChatUsername(username string) *unbanChatMember {
	ucm.chatId = username
	return ucm
}

/*func (ucm *unbanChatMember) Send() (message *structs.Message, err error) {
	message, err = ucm.parent.Send(sd)
	return
}
*/
