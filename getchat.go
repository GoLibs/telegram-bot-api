package go_telegram_bot_api

import (
	"encoding/json"

	"github.com/golibs/telegram-bot-api/structs"
)

type getChat struct {
	parent *TelegramBot
	chatId interface{}
}

func (sv *getChat) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
	}{
		ChatId: sv.chatId,
	})
}

func (sv *getChat) response() interface{} {
	var result structs.Chat
	return &result
}

func (sv *getChat) method() string {
	return "POST"
}

func (sv *getChat) endpoint() string {
	return "getChat"
}

func (sv *getChat) SetChatId(chatId int64) *getChat {
	sv.chatId = chatId
	return sv
}

func (sv *getChat) SetChatUsername(username string) *getChat {
	sv.chatId = username
	return sv
}

/*func (sv *getChat) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
