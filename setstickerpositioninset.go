package go_telegram_bot_api

import (
	"encoding/json"
)

type setStickerPositionInSet struct {
	parent   *TelegramBot
	sticker  string
	position int64
}

func (sv *setStickerPositionInSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sticker  string `json:"sticker"`
		Position int64  `json:"position"`
	}{
		Sticker:  sv.sticker,
		Position: sv.position,
	})
}

func (sv *setStickerPositionInSet) response() interface{} {
	var resp bool
	return &resp
}

func (sv *setStickerPositionInSet) method() string {
	return "POST"
}

func (sv *setStickerPositionInSet) endpoint() string {
	return "setStickerPositionInSet"
}

func (sv *setStickerPositionInSet) SetSticker(sticker string) *setStickerPositionInSet {
	sv.sticker = sticker
	return sv
}

func (sv *setStickerPositionInSet) SetPosition(pos int64) *setStickerPositionInSet {
	sv.position = pos
	return sv
}
