package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendMessage(t *testing.T) {
	/*var text tgbotapi.Text
	tex := text.AddWithNewLine("Hello").AddWithNewLine("Boy")*/
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
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
	fmt.Println(m)
}
