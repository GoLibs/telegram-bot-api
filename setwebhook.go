package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type setWebhook struct {
	parent *TelegramBot
	url    string
	// certificate interface{} TODO
	iPAddress          string
	maxConnections     int
	allowedUpdates     []string
	dropPendingUpdates bool
}

func (sw *setWebhook) SetUrl(address string) *setWebhook {
	sw.url = address
	return sw
}

func (sw *setWebhook) SetIPAddress(address string) *setWebhook {
	sw.iPAddress = address
	return sw
}

func (sw *setWebhook) SetMaxConnections(max int) *setWebhook {
	sw.maxConnections = max
	return sw
}

func (sw *setWebhook) DropPendingUpdates() *setWebhook {
	sw.dropPendingUpdates = true
	return sw
}

func (sw *setWebhook) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Url string `json:"url,omitempty"`
		// certificate interface{} TODO
		IPAddress          string   `json:"ip_address,omitempty"`
		MaxConnections     int      `json:"max_connections,omitempty"`
		AllowedUpdates     []string `json:"allowed_updates,omitempty"`
		DropPendingUpdates bool     `json:"drop_pending_updates,omitempty"`
	}{
		Url:                sw.url,
		IPAddress:          sw.iPAddress,
		MaxConnections:     sw.maxConnections,
		AllowedUpdates:     sw.allowedUpdates,
		DropPendingUpdates: sw.dropPendingUpdates,
	})
}

func (sw *setWebhook) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sw *setWebhook) method() string {
	return "POST"
}

func (sw *setWebhook) endpoint() string {
	return "setWebhook"
}

func (sw *setWebhook) Set() (res bool, err error) {
	var result *Response
	result, err = sw.parent.Send(sw)
	if err != nil {
		return
	}
	res = *result.Bool
	return
}
