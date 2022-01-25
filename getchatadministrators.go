package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type getChatAdministrators struct {
	parent *TelegramBot
	chatId interface{}
}

func (sv *getChatAdministrators) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
	}{
		ChatId: sv.chatId,
	})
}

func (sv *getChatAdministrators) response() interface{} {
	var result []structs.ChatMember
	return &result
}

func (sv *getChatAdministrators) method() string {
	return "POST"
}

func (sv *getChatAdministrators) endpoint() string {
	return "getChatAdministrators"
}

func (sv *getChatAdministrators) SetChatId(chatId int64) *getChatAdministrators {
	sv.chatId = chatId
	return sv
}

func (sv *getChatAdministrators) SetChatUsername(username string) *getChatAdministrators {
	sv.chatId = username
	return sv
}

/*func (sv *getChatAdministrators) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
