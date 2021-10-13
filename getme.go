package go_telegram_bot_api

import (
	"github.com/GoLibs/telegram-bot-api/structs"
)

type getMe struct {
	parent *TelegramBot
}

func (sv *getMe) marshalJSON() ([]byte, error) {
	return nil, nil
}

func (sv *getMe) response() interface{} {
	return &structs.User{}
}

func (sv *getMe) method() string {
	return "POST"
}

func (sv *getMe) endpoint() string {
	return "getMe"
}

/*func (sv *getMe) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
