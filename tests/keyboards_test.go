package tests

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestKeyboards(t *testing.T) {
	/**/
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Inline Keyboard Test */
	keyboard := tgbotapi.TelegramBot{}.Tools.Keyboards.NewInlineKeyboard()
	row := (*keyboard).NewRow()
	row.AddCallbackButton("hello", "set_hello")
	message := (bot.Message()).SetChatId(Tests{}.Defaults().UserId).SetText("good idea").SetReplyMarkup(keyboard)
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err, "err")
		return
	}
	pretty.Println(m)
	keyboard = tgbotapi.TelegramBot{}.Tools.Keyboards.NewInlineKeyboardFromSlicesOfMaps([][]map[string]string{
		{
			{
				"text":          "hello",
				"callback_data": "set_hello",
			},
		},
	})
	message.SetReplyMarkup(keyboard)
	m, err = bot.Send(message)
	if err != nil {
		fmt.Println(err, "err")
		return
	}
	pretty.Println(m)
	/* Reply Keyboard Test */
	message.SetReplyMarkup(bot.Tools.Keyboards.NewReplyKeyboardFromSlicesOfStrings([][]string{
		{"Hello Boy", "Buy"},
		{"See Ya", "Soon"},
		{"Well", "Played"},
	}))

	m, err = bot.Send(message)
	if err != nil {
		fmt.Println(err, "err")
		return
	}
	pretty.Println(m)
}
