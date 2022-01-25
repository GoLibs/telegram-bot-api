package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type setChatTitle struct {
	parent *TelegramBot
	chatId interface{}
	title  string
}

func (sv *setChatTitle) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
		Title  string      `json:"title"`
	}{
		ChatId: sv.chatId,
		Title:  sv.title,
	})
}

func (sv *setChatTitle) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *setChatTitle) method() string {
	return "POST"
}

func (sv *setChatTitle) endpoint() string {
	return "setChatTitle"
}

func (sv *setChatTitle) SetChatId(chatId int64) *setChatTitle {
	sv.chatId = chatId
	return sv
}

func (sv *setChatTitle) SetChatUsername(username string) *setChatTitle {
	sv.chatId = username
	return sv
}

func (sv *setChatTitle) SetTitle(title string) *setChatTitle {
	sv.title = title
	return sv
}

/*func (sv *setChatTitle) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
