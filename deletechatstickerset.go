package go_telegram_bot_api

import (
	"encoding/json"
)

type deleteChatStickerSet struct {
	parent *TelegramBot
	chatId interface{}
}

func (sv *deleteChatStickerSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
	}{
		ChatId: sv.chatId,
	})
}

func (sv *deleteChatStickerSet) response() interface{} {
	var result bool
	return &result
}

func (sv *deleteChatStickerSet) method() string {
	return "POST"
}

func (sv *deleteChatStickerSet) endpoint() string {
	return "deleteChatStickerSet"
}

func (sv *deleteChatStickerSet) SetChatId(chatId int64) *deleteChatStickerSet {
	sv.chatId = chatId
	return sv
}

func (sv *deleteChatStickerSet) SetChatUsername(username string) *deleteChatStickerSet {
	sv.chatId = username
	return sv
}

/*func (sv *deleteChatStickerSet) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
