package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestCopyMessage(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
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
