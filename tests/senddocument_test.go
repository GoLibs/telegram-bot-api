package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api/structs"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendDocument(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	k := structs.ReplyKeyboardMarkup{}.FromSlicesOfStrings([][]string{{"g", "h"}, {"s", "j"}})
	message := (bot.Document()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetDocumentFilePath("photo.png").SetReplyMarkup(k)
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Document)

	message = (bot.Document()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetDocumentId("BQACAgQAAxkDAAECEwZeYXI771gPsruqrYEjhVewsKz8qAACJgYAAiWdEVPMjo7I4Zz4mxgE")
	m, err = bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Document)

	/*	m, err = (bot.Document()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetDocumentThumbFilePath("photo.png").SetDocumentId("BQACAgQAAxkDAAECEwZeYXI771gPsruqrYEjhVewsKz8qAACJgYAAiWdEVPMjo7I4Zz4mxgE").Send()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*m.Message.Document)*/
}
