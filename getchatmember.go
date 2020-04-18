package go_telegram_bot_api

import (
	"encoding/json"

	"github.com/golibs/telegram-bot-api/structs"
)

type getChatMember struct {
	parent *TelegramBot
	chatId interface{}
	userId int64
}

func (sv *getChatMember) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
		UserId int64       `json:"user_id"`
	}{
		ChatId: sv.chatId,
		UserId: sv.userId,
	})
}

func (sv *getChatMember) response() interface{} {
	var result structs.ChatMember
	return &result
}

func (sv *getChatMember) method() string {
	return "POST"
}

func (sv *getChatMember) endpoint() string {
	return "getChatAdministrators"
}

func (sv *getChatMember) SetChatId(chatId int64) *getChatMember {
	sv.chatId = chatId
	return sv
}

func (sv *getChatMember) SetChatUsername(username string) *getChatMember {
	sv.chatId = username
	return sv
}

func (sv *getChatMember) SetUserId(userId int64) *getChatMember {
	sv.userId = userId
	return sv
}

/*func (sv *getChatMember) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
