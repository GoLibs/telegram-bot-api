package go_telegram_bot_api

import (
	"encoding/json"

	"github.com/golibs/telegram-bot-api/structs"
)

type getUserProfilePhotos struct {
	parent *TelegramBot
	userId int64
	offset int64
	limit  int64
}

func (sv *getUserProfilePhotos) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId int64 `json:"user_id"`
		Offset int64 `json:"offset,omitempty"`
		Limit  int64 `json:"limit,omitempty"`
	}{
		UserId: sv.userId,
		Offset: sv.offset,
		Limit:  sv.limit,
	})
}

func (sv *getUserProfilePhotos) response() interface{} {
	return &structs.UserProfilePhotos{}
}

func (sv *getUserProfilePhotos) method() string {
	return "POST"
}

func (sv *getUserProfilePhotos) endpoint() string {
	return "getUserProfilePhotos"
}

func (sv *getUserProfilePhotos) SetUserId(userId int64) *getUserProfilePhotos {
	sv.userId = userId
	return sv
}

func (sv *getUserProfilePhotos) SetOffset(offset int64) *getUserProfilePhotos {
	sv.offset = offset
	return sv
}

func (sv *getUserProfilePhotos) SetLimit(limit int64) *getUserProfilePhotos {
	sv.limit = limit
	return sv
}

/*func (sv *getUserProfilePhotos) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
