package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type getStickerSet struct {
	parent *TelegramBot
	name   string
}

func (sv *getStickerSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name string `json:"name"`
	}{
		Name: sv.name,
	})
}

func (sv *getStickerSet) response() interface{} {
	return &structs.StickerSet{}
}

func (sv *getStickerSet) method() string {
	return "POST"
}

func (sv *getStickerSet) endpoint() string {
	return "getStickerSet"
}

func (sv *getStickerSet) SetName(name string) *getStickerSet {
	sv.name = name
	return sv
}

/*func (sv *getStickerSet) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
