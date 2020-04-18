package go_telegram_bot_api

import (
	"encoding/json"
)

type getChatMembersCount struct {
	parent *TelegramBot
	chatId interface{}
}

func (sv *getChatMembersCount) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
	}{
		ChatId: sv.chatId,
	})
}

func (sv *getChatMembersCount) response() interface{} {
	var result int64
	return &result
}

func (sv *getChatMembersCount) method() string {
	return "POST"
}

func (sv *getChatMembersCount) endpoint() string {
	return "getChat"
}

func (sv *getChatMembersCount) SetChatId(chatId int64) *getChatMembersCount {
	sv.chatId = chatId
	return sv
}

func (sv *getChatMembersCount) SetChatUsername(username string) *getChatMembersCount {
	sv.chatId = username
	return sv
}

/*func (sv *getChatMembersCount) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
