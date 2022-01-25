package tgbotapi

import (
	"encoding/json"
	"io"
	"time"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type addStickerToSet struct {
	parent       *TelegramBot
	userId       int64
	name         string
	pngSticker   interface{}
	emojis       string
	maskPosition *structs.MaskPosition
	pngFile      *fileInfo
}

func (sv *addStickerToSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId       int64                 `json:"user_id"`
		Name         string                `json:"name"`
		PngSticker   interface{}           `json:"png_sticker"`
		Emojis       string                `json:"emojis"`
		MaskPosition *structs.MaskPosition `json:"mask_position"`
	}{
		UserId:       sv.userId,
		Name:         sv.name,
		PngSticker:   sv.pngSticker,
		Emojis:       sv.emojis,
		MaskPosition: sv.maskPosition,
	})
}

func (sv *addStickerToSet) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *addStickerToSet) method() string {
	return "POST"
}

func (sv *addStickerToSet) endpoint() string {
	return "addStickerToSet"
}

func (sv *addStickerToSet) medias() []fileInfo {
	if sv.pngFile != nil {
		return []fileInfo{*sv.pngFile}
	}
	return nil
}

func (sv *addStickerToSet) SetUserId(userId int64) *addStickerToSet {
	sv.userId = userId
	return sv
}

func (sv *addStickerToSet) SetName(name string) *addStickerToSet {
	sv.name = name
	return sv
}

func (sv *addStickerToSet) SetEmojis(emojis string) *addStickerToSet {
	sv.emojis = emojis
	return sv
}

/*func (sv *addStickerToSet) SetMaskPosition(mask *structs.MaskPosition) *addStickerToSet {
	sv.maskPosition = mask
	return sv
}*/

func (sv *addStickerToSet) SetPngStickerId(pngStickerId string) *addStickerToSet {
	sv.pngSticker = pngStickerId
	return sv
}

func (sv *addStickerToSet) SetPngStickerFilePath(stickerFilePath string) *addStickerToSet {
	sv.pngFile = &fileInfo{
		Field: "sticker",
		Path:  stickerFilePath,
	}
	sv.pngSticker = "attach://sticker"
	return sv
}

func (sv *addStickerToSet) SetPngStickerFileReader(stickerFileReader io.Reader, fileName string) *addStickerToSet {
	if fileName == "" {
		fileName = time.Now().Format("2006_01_02_15_04_05")
	}
	sv.pngFile = &fileInfo{
		Field:  "sticker",
		Reader: stickerFileReader,
		Name:   fileName,
	}
	sv.pngSticker = "attach://sticker"
	return sv
}
