package tgbotapi

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type deleteWebhook struct {
	parent             *TelegramBot
	dropPendingUpdates bool
}

func (dw *deleteWebhook) DropPendingUpdates() *deleteWebhook {
	dw.dropPendingUpdates = true
	return dw
}

func (dw *deleteWebhook) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
	}{
		DropPendingUpdates: dw.dropPendingUpdates,
	})
}

func (dw *deleteWebhook) response() interface{} {
	return structs.ResponseTypeBool()
}

func (dw *deleteWebhook) method() string {
	return "POST"
}

func (dw *deleteWebhook) endpoint() string {
	return "deleteWebhook"
}

func (dw *deleteWebhook) Set() (res bool, err error) {
	var result *Response
	result, err = dw.parent.Send(dw)
	if err != nil {
		return
	}
	res = *result.Bool
	return
}
