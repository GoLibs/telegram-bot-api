package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestSendVenue(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	loc := (bot.Venue()).SetChatId(Tests{}.Defaults().UserId).SetLatitude(20.073750).SetLongitude(30.280556)
	resp, err := bot.Send(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Message)
}
