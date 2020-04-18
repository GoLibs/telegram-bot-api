package tests

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestSendMessage(t *testing.T) {
	/*var text go_telegram_bot_api.Text
	tex := text.AddWithNewLine("Hello").AddWithNewLine("Boy")*/
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	message := (bot.Message()).SetChatId(Tests{}.Defaults().UserId).SetText("good idea" + bot.Tools.Strings.EmptyNewLine())
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err, "err")
		return
	}
	pretty.Println(m)
}
