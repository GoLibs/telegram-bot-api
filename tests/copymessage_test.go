package tests

import (
	"fmt"
	"testing"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestCopyMessage(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := bot.Send(bot.Message().SetChatId(Tests{}.Defaults().UserId).SetText("hi"))
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err = bot.Send(bot.CopyMessage().SetMessage(res.Message).SetChatId(Tests{}.Defaults().UserId))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.MessageId)
}
