package go_telegram_bot_api

import (
	"encoding/json"
)

type setChatDescription struct {
	parent      *TelegramBot
	chatId      interface{}
	description string
}

func (sv *setChatDescription) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId      interface{} `json:"chat_id"`
		Description string      `json:"description"`
	}{
		ChatId:      sv.chatId,
		Description: sv.description,
	})
}

func (sv *setChatDescription) response() interface{} {
	var result bool
	return &result
}

func (sv *setChatDescription) method() string {
	return "POST"
}

func (sv *setChatDescription) endpoint() string {
	return "setChatDescription"
}

func (sv *setChatDescription) SetChatId(chatId int64) *setChatDescription {
	sv.chatId = chatId
	return sv
}

func (sv *setChatDescription) SetChatUsername(username string) *setChatDescription {
	sv.chatId = username
	return sv
}

func (sv *setChatDescription) SetDescription(description string) *setChatDescription {
	sv.description = description
	return sv
}

/*func (sv *setChatDescription) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
