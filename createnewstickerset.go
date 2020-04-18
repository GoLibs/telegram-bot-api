package go_telegram_bot_api

import (
	"encoding/json"
	"io"
	"time"

	"github.com/golibs/telegram-bot-api/structs"
)

type createNewStickerSet struct {
	parent        *TelegramBot
	userId        int64
	name          string
	title         string
	pngSticker    interface{}
	emojis        string
	containsMasks bool
	maskPosition  *structs.MaskPosition
	pngFile       *fileInfo
}

func (sv *createNewStickerSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId        int64                 `json:"user_id"`
		Name          string                `json:"name"`
		Title         string                `json:"title"`
		PngSticker    interface{}           `json:"png_sticker"`
		Emojis        string                `json:"emojis"`
		ContainsMasks bool                  `json:"contains_masks"`
		MaskPosition  *structs.MaskPosition `json:"mask_position"`
	}{
		UserId:        sv.userId,
		Name:          sv.name,
		Title:         sv.title,
		PngSticker:    sv.pngSticker,
		Emojis:        sv.emojis,
		ContainsMasks: sv.containsMasks,
		MaskPosition:  sv.maskPosition,
	})
}

func (sv *createNewStickerSet) response() interface{} {
	var resp bool
	return &resp
}

func (sv *createNewStickerSet) method() string {
	return "POST"
}

func (sv *createNewStickerSet) endpoint() string {
	return "createNewStickerSet"
}

func (sv *createNewStickerSet) medias() []fileInfo {
	if sv.pngFile != nil {
		return []fileInfo{*sv.pngFile}
	}
	return nil
}

func (sv *createNewStickerSet) SetUserId(userId int64) *createNewStickerSet {
	sv.userId = userId
	return sv
}

func (sv *createNewStickerSet) SetName(name string) *createNewStickerSet {
	sv.name = name
	return sv
}

func (sv *createNewStickerSet) SetTitle(title string) *createNewStickerSet {
	sv.title = title
	return sv
}

func (sv *createNewStickerSet) SetEmojis(emojis string) *createNewStickerSet {
	sv.emojis = emojis
	return sv
}

func (sv *createNewStickerSet) SetContainsMasks(b bool) *createNewStickerSet {
	sv.containsMasks = b
	return sv
}

/*func (sv *createNewStickerSet) SetMaskPosition(mask *structs.MaskPosition) *createNewStickerSet {
	sv.maskPosition = mask
	return sv
}*/

func (sv *createNewStickerSet) SetPngStickerId(pngStickerId string) *createNewStickerSet {
	sv.pngSticker = pngStickerId
	return sv
}

func (sv *createNewStickerSet) SetPngStickerFilePath(stickerFilePath string) *createNewStickerSet {
	sv.pngFile = &fileInfo{
		Field: "sticker",
		Path:  stickerFilePath,
	}
	sv.pngSticker = "attach://sticker"
	return sv
}

func (sv *createNewStickerSet) SetPngStickerFileReader(stickerFileReader io.Reader, fileName string) *createNewStickerSet {
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
