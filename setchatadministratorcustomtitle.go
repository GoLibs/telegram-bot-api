package go_telegram_bot_api

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type setChatAdministratorCustomTitle struct {
	parent      *TelegramBot
	chatId      interface{}
	userId      int64
	customTitle string
}

func (sv *setChatAdministratorCustomTitle) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId      interface{} `json:"chat_id"`
		UserId      int64       `json:"user_id"`
		CustomTitle string      `json:"custom_title"`
	}{
		ChatId:      sv.chatId,
		UserId:      sv.userId,
		CustomTitle: sv.customTitle,
	})
}

func (sv *setChatAdministratorCustomTitle) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *setChatAdministratorCustomTitle) method() string {
	return "POST"
}

func (sv *setChatAdministratorCustomTitle) endpoint() string {
	return "setChatAdministratorCustomTitle"
}

func (sv *setChatAdministratorCustomTitle) SetChatId(chatId int64) *setChatAdministratorCustomTitle {
	sv.chatId = chatId
	return sv
}

func (sv *setChatAdministratorCustomTitle) SetChatUsername(username string) *setChatAdministratorCustomTitle {
	sv.chatId = username
	return sv
}

func (sv *setChatAdministratorCustomTitle) SetUserId(userId int64) *setChatAdministratorCustomTitle {
	sv.userId = userId
	return sv
}

func (sv *setChatAdministratorCustomTitle) SetCustomTitle(title string) *setChatAdministratorCustomTitle {
	sv.customTitle = title
	return sv
}

/*func (sv *setChatAdministratorCustomTitle) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
