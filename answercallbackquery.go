package tgbotapi

import (
	"encoding/json"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type answerCallbackQuery struct {
	parent          *TelegramBot
	callbackQueryId string
	text            string
	showAlert       bool
	url             string
	cacheTime       int64
}

func (sv *answerCallbackQuery) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		CallbackQueryId string `json:"callback_query_id"`
		Text            string `json:"text,omitempty"`
		ShowAlert       bool   `json:"show_alert,omitempty"`
		Url             string `json:"url,omitempty"`
		CacheTime       int64  `json:"cache_time,omitempty"`
	}{
		CallbackQueryId: sv.callbackQueryId,
		Text:            sv.text,
		ShowAlert:       sv.showAlert,
		Url:             sv.url,
		CacheTime:       sv.cacheTime,
	})
}

func (sv *answerCallbackQuery) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *answerCallbackQuery) method() string {
	return "POST"
}

func (sv *answerCallbackQuery) endpoint() string {
	return "answerCallbackQuery"
}

func (sv *answerCallbackQuery) SetCallbackQueryId(callbackQueryId string) *answerCallbackQuery {
	sv.callbackQueryId = callbackQueryId
	return sv
}

func (sv *answerCallbackQuery) SetText(text string) *answerCallbackQuery {
	sv.text = text
	return sv
}

func (sv *answerCallbackQuery) SetShowAlert(showAlert bool) *answerCallbackQuery {
	sv.showAlert = showAlert
	return sv
}

func (sv *answerCallbackQuery) SetUrl(url string) *answerCallbackQuery {
	sv.url = url
	return sv
}

func (sv *answerCallbackQuery) SetCacheTime(cacheTime int64) *answerCallbackQuery {
	sv.cacheTime = cacheTime
	return sv
}
