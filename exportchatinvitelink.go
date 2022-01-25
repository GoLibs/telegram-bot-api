package tgbotapi

import (
	"encoding/json"
)

type exportChatInviteLink struct {
	parent *TelegramBot
	chatId interface{}
}

func (sv *exportChatInviteLink) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
	}{
		ChatId: sv.chatId,
	})
}

func (sv *exportChatInviteLink) response() interface{} {
	var result string
	return &result
}

func (sv *exportChatInviteLink) method() string {
	return "POST"
}

func (sv *exportChatInviteLink) endpoint() string {
	return "exportChatInviteLink"
}

func (sv *exportChatInviteLink) SetChatId(chatId int64) *exportChatInviteLink {
	sv.chatId = chatId
	return sv
}

func (sv *exportChatInviteLink) SetChatUsername(username string) *exportChatInviteLink {
	sv.chatId = username
	return sv
}

/*func (sv *exportChatInviteLink) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
