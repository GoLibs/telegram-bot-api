package tgbotapi

import (
	"encoding/json"
	"io"
)

type setChatPhoto struct {
	parent *TelegramBot
	chatId interface{}
	photo  interface{}
	file
	photos []fileInfo
}

func (sv *setChatPhoto) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
		Photo  interface{} `json:"photo,omitempty"`
	}{
		ChatId: sv.chatId,
		Photo:  sv.photo,
	})
}

func (sv *setChatPhoto) response() interface{} {
	var result string
	return &result
}

func (sv *setChatPhoto) medias() []fileInfo {
	for i := range sv.photos {
		sv.photos[i].Field = "photo"
	}
	return sv.photos
}

func (sv *setChatPhoto) method() string {
	return "POST"
}

func (sv *setChatPhoto) endpoint() string {
	return "exportChatInviteLink"
}

func (sv *setChatPhoto) SetChatId(chatId int64) *setChatPhoto {
	sv.chatId = chatId
	return sv
}

func (sv *setChatPhoto) SetChatUsername(username string) *setChatPhoto {
	sv.chatId = username
	return sv
}

func (sv *setChatPhoto) SetPhotoId(photoId string) *setChatPhoto {
	sv.photo = photoId
	return sv
}

func (sv *setChatPhoto) SetPhotoFilePath(photoFilePath string) *setChatPhoto {
	if sv.photos == nil {
		sv.photos = []fileInfo{}
	}
	sv.photos = append(sv.photos, fileInfo{
		Path: photoFilePath,
	})
	return sv
}

func (sv *setChatPhoto) SetPhotoReader(phr io.Reader, photoName string) *setChatPhoto {
	if sv.photos == nil {
		sv.photos = []fileInfo{}
	}
	sv.photos = append(sv.photos, fileInfo{
		Reader: phr,
		Name:   photoName,
	})
	return sv
}

/*func (sv *setChatPhoto) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
