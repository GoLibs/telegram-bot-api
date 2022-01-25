package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type deleteChatPhoto struct {
	parent *TelegramBot
	chatId interface{}
	photo  interface{}
}

func (sv *deleteChatPhoto) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
		Photo  interface{} `json:"photo,omitempty"`
	}{
		ChatId: sv.chatId,
		Photo:  sv.photo,
	})
}

func (sv *deleteChatPhoto) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *deleteChatPhoto) method() string {
	return "POST"
}

func (sv *deleteChatPhoto) endpoint() string {
	return "deleteChatPhoto"
}

func (sv *deleteChatPhoto) SetChatId(chatId int64) *deleteChatPhoto {
	sv.chatId = chatId
	return sv
}

func (sv *deleteChatPhoto) SetChatUsername(username string) *deleteChatPhoto {
	sv.chatId = username
	return sv
}
