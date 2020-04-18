package go_telegram_bot_api

import (
	"encoding/json"
)

type deleteStickerFromSet struct {
	parent  *TelegramBot
	sticker string
}

func (sv *deleteStickerFromSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sticker  string `json:"sticker"`
		Position int64  `json:"position"`
	}{
		Sticker: sv.sticker,
	})
}

func (sv *deleteStickerFromSet) response() interface{} {
	var resp bool
	return &resp
}

func (sv *deleteStickerFromSet) method() string {
	return "POST"
}

func (sv *deleteStickerFromSet) endpoint() string {
	return "deleteStickerFromSet"
}

func (sv *deleteStickerFromSet) SetSticker(sticker string) *deleteStickerFromSet {
	sv.sticker = sticker
	return sv
}
