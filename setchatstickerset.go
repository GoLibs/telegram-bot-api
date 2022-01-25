package tgbotapi

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type setChatStickerSet struct {
	parent         *TelegramBot
	chatId         interface{}
	stickerSetName string
}

func (sv *setChatStickerSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId         interface{} `json:"chat_id"`
		StickerSetName string      `json:"sticker_set_name"`
	}{
		ChatId:         sv.chatId,
		StickerSetName: sv.stickerSetName,
	})
}

func (sv *setChatStickerSet) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *setChatStickerSet) method() string {
	return "POST"
}

func (sv *setChatStickerSet) endpoint() string {
	return "setChatStickerSet"
}

func (sv *setChatStickerSet) SetChatId(chatId int64) *setChatStickerSet {
	sv.chatId = chatId
	return sv
}

func (sv *setChatStickerSet) SetChatUsername(username string) *setChatStickerSet {
	sv.chatId = username
	return sv
}

func (sv *setChatStickerSet) SetStickerSetName(name string) *setChatStickerSet {
	sv.stickerSetName = name
	return sv
}

/*func (sv *setChatStickerSet) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
