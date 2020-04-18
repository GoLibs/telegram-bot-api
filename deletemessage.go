package go_telegram_bot_api

import (
	"encoding/json"
)

type deleteMessage struct {
	parent    *TelegramBot
	chatId    interface{}
	messageId int64
}

func (m *deleteMessage) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId    interface{} `json:"chat_id"`
		MessageId int64       `json:"message_id"`
	}{
		ChatId:    m.chatId,
		MessageId: m.messageId,
	})
}

func (m *deleteMessage) response() interface{} {
	var resp bool
	return &resp
}

func (m *deleteMessage) method() string {
	return "POST"
}

func (m *deleteMessage) endpoint() string {
	return "deleteMessage"
}

func (m *deleteMessage) SetChatId(chatId int64) *deleteMessage {
	m.chatId = chatId
	return m
}

func (m *deleteMessage) SetChatUsername(username string) *deleteMessage {
	m.chatId = username
	return m
}

func (m *deleteMessage) SetMessageId(messageId int64) *deleteMessage {
	m.messageId = messageId
	return m
}
