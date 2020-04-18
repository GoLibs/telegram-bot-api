package structs

import "encoding/json"

type InlineKeyboardButton struct {
	text                         string
	url                          string
	loginUrl                     LoginUrl
	callbackData                 string
	switchInlineQuery            string
	switchInlineQueryCurrentChat string
	callbackGame                 interface{}
	pay                          bool
}

func (i InlineKeyboardButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Text                         string      `json:"text"`
		Url                          string      `json:"url"`
		LoginUrl                     LoginUrl    `json:"login_url"`
		CallbackData                 string      `json:"callback_data"`
		SwitchInlineQuery            string      `json:"switch_inline_query"`
		SwitchInlineQueryCurrentChat string      `json:"switch_inline_query_current_chat"`
		CallbackGame                 interface{} `json:"callback_game"`
		Pay                          bool        `json:"pay"`
	}{
		Text:                         i.text,
		Url:                          i.url,
		LoginUrl:                     i.loginUrl,
		CallbackData:                 i.callbackData,
		SwitchInlineQuery:            i.switchInlineQuery,
		SwitchInlineQueryCurrentChat: i.switchInlineQueryCurrentChat,
		CallbackGame:                 i.callbackGame,
		Pay:                          i.pay,
	})
}

func (i *InlineKeyboardButton) SetText(text string) {
	(*i).text = text
}

//
// func (i *InlineKeyboardButton) SetUrl(url string) {
// 	i.url = url
// }
//
// func (i *InlineKeyboardButton) SetLoginUrl(loginUrl LoginUrl) {
// 	i.loginUrl = loginUrl
// }
//
// func (i *InlineKeyboardButton) SetCallbackData(callbackData string) {
// 	i.callbackData = callbackData
// }
//
// func (i *InlineKeyboardButton) SetSwitchInlineQuery(switchInlineQuery string) {
// 	i.switchInlineQuery = switchInlineQuery
// }
//
// func (i *InlineKeyboardButton) SwitchInlineQueryCurrentChat() string {
// 	return i.switchInlineQueryCurrentChat
// }
//
// func (i *InlineKeyboardButton) SetCallbackGame(callbackGame interface{}) {
// 	i.callbackGame = callbackGame
// }
//
// func (i *InlineKeyboardButton) EnablePay() {
// 	i.pay = true
// }
//
// func (i *InlineKeyboardButton) DisablePay() {
// 	i.pay = false
// }
