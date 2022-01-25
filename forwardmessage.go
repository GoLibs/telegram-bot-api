package tgbotapi

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type forwardMessage struct {
	parent              *TelegramBot
	chatId              interface{}
	fromChatId          interface{}
	disableNotification bool
	messageId           int64
}

func (fm *forwardMessage) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		FromChatId          interface{} `json:"from_chat_id"`
		DisableNotification bool        `json:"disable_notification"`
		MessageId           int64       `json:"message_id"`
	}{
		ChatId:              fm.chatId,
		FromChatId:          fm.fromChatId,
		DisableNotification: fm.disableNotification,
		MessageId:           fm.messageId,
	})
}

func (fm *forwardMessage) response() interface{} {
	return &structs.Message{}
}

func (fm *forwardMessage) SetChatId(chatId int64) *forwardMessage {
	fm.chatId = chatId
	return fm
}

func (fm *forwardMessage) SetChatUsername(username string) *forwardMessage {
	fm.chatId = username
	return fm
}

func (fm *forwardMessage) SetFromChatId(fromChatId int64) *forwardMessage {
	fm.fromChatId = fromChatId
	return fm
}

func (fm *forwardMessage) SetFromChatUsername(fromChatUsername string) *forwardMessage {
	fm.fromChatId = fromChatUsername
	return fm
}

func (fm *forwardMessage) SetDisableNotification() *forwardMessage {
	fm.disableNotification = true
	return fm
}

func (fm *forwardMessage) SetEnableNotification() *forwardMessage {
	fm.disableNotification = false
	return fm
}

func (fm *forwardMessage) SetMessageId(messageId int64) *forwardMessage {
	fm.messageId = messageId
	return fm
}

func (fm *forwardMessage) SetMessage(message *structs.Message) *forwardMessage {
	fm.messageId = message.MessageId
	fm.fromChatId = message.Chat.Id
	return fm
}

func (fm *forwardMessage) method() string {
	return "POST"
}

func (fm *forwardMessage) endpoint() string {
	return "forwardMessage"
}

/*func (fm *forwardMessage) Forward() (message *structs.Message, err error) {
	message, err = fm.parent.Send(fm)
	return
}
*/
