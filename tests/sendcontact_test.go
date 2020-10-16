package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestSendContact(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	loc := (bot.Contact()).SetChatId(Tests{}.Defaults().UserId).SetPhoneNumber("+132352123").SetFirstName("Ali")
	resp, err := bot.Send(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Message)
}
