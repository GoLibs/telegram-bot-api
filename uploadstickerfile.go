package go_telegram_bot_api

import (
	"encoding/json"
	"io"
	"time"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type uploadStickerFile struct {
	parent     *TelegramBot
	userId     int64
	pngSticker interface{}
	file
	fileInfo *fileInfo
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sv *uploadStickerFile) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId     int64       `json:"user_id"`
		PngSticker interface{} `json:"png_sticker"`
	}{
		UserId:     sv.userId,
		PngSticker: sv.pngSticker,
	})
}

func (sv *uploadStickerFile) response() interface{} {
	return &structs.File{}
}

func (sv *uploadStickerFile) method() string {
	return "POST"
}

func (sv *uploadStickerFile) endpoint() string {
	return "uploadStickerFile"
}

func (sv *uploadStickerFile) medias() []fileInfo {
	if sv.fileInfo != nil {
		return []fileInfo{*sv.fileInfo}
	}
	return nil
}

func (sv *uploadStickerFile) SetUserId(userId int64) *uploadStickerFile {
	sv.userId = userId
	return sv
}

func (sv *uploadStickerFile) SetStickerFilePath(stickerFilePath string) *uploadStickerFile {
	sv.fileInfo = &fileInfo{
		Field: "sticker",
		Path:  stickerFilePath,
	}
	sv.pngSticker = "attach://sticker"
	return sv
}

func (sv *uploadStickerFile) SetStickerFileReader(stickerFileReader io.Reader, fileName string) *uploadStickerFile {
	if fileName == "" {
		fileName = time.Now().Format("2006_01_02_15_04_05")
	}
	sv.fileInfo = &fileInfo{
		Field:  "sticker",
		Reader: stickerFileReader,
		Name:   fileName,
	}
	sv.pngSticker = "attach://sticker"
	return sv
}

/*func (sv *uploadStickerFile) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
