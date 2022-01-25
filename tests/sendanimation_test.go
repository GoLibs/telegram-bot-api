package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendAnimation(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := (bot.Animation()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAnimationFilePath("animation.mp4")
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Animation)

	message = (bot.Animation()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAnimationId("CgACAgQAAxkDAAECEwpeYX4Q_c5R9uOxZwJKqpq_6Usc8QACKgYAAiWdEVOpN11cGifpehgE")
	m, err = bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Animation)
}
