package go_telegram_bot_api

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type unpinChatMessage struct {
	parent *TelegramBot
	chatId interface{}
}

func (sv *unpinChatMessage) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
	}{
		ChatId: sv.chatId,
	})
}

func (sv *unpinChatMessage) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *unpinChatMessage) method() string {
	return "POST"
}

func (sv *unpinChatMessage) endpoint() string {
	return "unpinChatMessage"
}

func (sv *unpinChatMessage) SetChatId(chatId int64) *unpinChatMessage {
	sv.chatId = chatId
	return sv
}

func (sv *unpinChatMessage) SetChatUsername(username string) *unpinChatMessage {
	sv.chatId = username
	return sv
}

/*func (sv *unpinChatMessage) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
