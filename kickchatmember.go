package go_telegram_bot_api

import (
	"encoding/json"
	"time"
)

type kickChatMember struct {
	parent    *TelegramBot
	chatId    interface{}
	userId    int64
	untilDate int64
}

func (kcm *kickChatMember) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId    interface{} `json:"chat_id"`
		UserId    int64       `json:"user_id"`
		UntilDate int64       `json:"until_date,omitempty"`
	}{
		ChatId:    kcm.chatId,
		UserId:    kcm.userId,
		UntilDate: kcm.untilDate,
	})
}

func (kcm *kickChatMember) response() interface{} {
	var resp bool
	return &resp
}

func (kcm *kickChatMember) method() string {
	return "POST"
}

func (kcm *kickChatMember) endpoint() string {
	return "kickChatMember"
}

func (kcm *kickChatMember) SetUserId(userId int64) *kickChatMember {
	kcm.userId = userId
	return kcm
}

func (kcm *kickChatMember) SetChatId(chatId int64) *kickChatMember {
	kcm.chatId = chatId
	return kcm
}

func (kcm *kickChatMember) SetChatUsername(username string) *kickChatMember {
	kcm.chatId = username
	return kcm
}

func (kcm *kickChatMember) SetUntilDate(untilDate int64) *kickChatMember {
	kcm.untilDate = untilDate
	return kcm
}

func (kcm *kickChatMember) SetPermanent() *kickChatMember {
	kcm.untilDate = time.Now().Add(-(time.Second * 31)).Unix()
	return kcm
}

/*func (kcm *kickChatMember) Send() (message *structs.Message, err error) {
	message, err = kcm.parent.Send(sd)
	return
}
*/
