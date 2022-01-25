package tgbotapi

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type sendChatAction struct {
	parent *TelegramBot
	chatId interface{}
	action string
}

func (sca *sendChatAction) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
		Action string      `json:"action"`
	}{
		ChatId: sca.chatId,
		Action: sca.action,
	})
}

func (sca *sendChatAction) response() interface{} {
	return &structs.Message{}
}

func (sca *sendChatAction) method() string {
	return "POST"
}

func (sca *sendChatAction) endpoint() string {
	return "sendContact"
}

func (sca *sendChatAction) SetIsTyping() *sendChatAction {
	sca.action = "typing"
	return sca
}

func (sca *sendChatAction) SetIsSendingPhoto() *sendChatAction {
	sca.action = "upload_photo"
	return sca
}

func (sca *sendChatAction) SetIsRecordingVideo() *sendChatAction {
	sca.action = "record_video"
	return sca
}

func (sca *sendChatAction) SetIsUploadingVideo() *sendChatAction {
	sca.action = "upload_video"
	return sca
}

func (sca *sendChatAction) SetIsRecordingVideoNote() *sendChatAction {
	sca.action = "record_video_note"
	return sca
}

func (sca *sendChatAction) SetIsUploadingVideoNote() *sendChatAction {
	sca.action = "upload_video_note"
	return sca
}

func (sca *sendChatAction) SetIsRecordingAudio() *sendChatAction {
	sca.action = "record_audio"
	return sca
}

func (sca *sendChatAction) SetIsUploadingAudio() *sendChatAction {
	sca.action = "upload_audio"
	return sca
}

func (sca *sendChatAction) SetIsUploadingDocument() *sendChatAction {
	sca.action = "upload_document"
	return sca
}

func (sca *sendChatAction) SetIsSendingLocation() *sendChatAction {
	sca.action = "find_location"
	return sca
}

func (sca *sendChatAction) SetChatUsername(username string) *sendChatAction {
	sca.chatId = username
	return sca
}

func (sca *sendChatAction) SetChatId(id int64) *sendChatAction {
	sca.chatId = id
	return sca
}

/*func (sca *sendChatAction) Send() (message *structs.Message, err error) {
	message, err = sca.parent.Send(sd)
	return
}
*/
