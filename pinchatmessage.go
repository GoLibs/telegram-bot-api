package go_telegram_bot_api

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type pinChatMessage struct {
	parent              *TelegramBot
	chatId              interface{}
	messageId           int64
	disableNotification bool
}

func (sv *pinChatMessage) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		MessageId           int64       `json:"message_id"`
		DisableNotification bool        `json:"disable_notification"`
	}{
		ChatId:              sv.chatId,
		MessageId:           sv.messageId,
		DisableNotification: sv.disableNotification,
	})
}

func (sv *pinChatMessage) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *pinChatMessage) method() string {
	return "POST"
}

func (sv *pinChatMessage) endpoint() string {
	return "pinChatMessage"
}

func (sv *pinChatMessage) SetChatId(chatId int64) *pinChatMessage {
	sv.chatId = chatId
	return sv
}

func (sv *pinChatMessage) SetChatUsername(username string) *pinChatMessage {
	sv.chatId = username
	return sv
}

func (sv *pinChatMessage) SetMessageId(messageId int64) *pinChatMessage {
	sv.messageId = messageId
	return sv
}

func (sv *pinChatMessage) EnableNotification() *pinChatMessage {
	sv.disableNotification = false
	return sv
}

func (sv *pinChatMessage) DisableNotification() *pinChatMessage {
	sv.disableNotification = true
	return sv
}

/*func (sv *pinChatMessage) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
