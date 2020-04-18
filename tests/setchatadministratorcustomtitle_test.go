package tests

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	go_telegram_bot_api "github.com/golibs/telegram-bot-api"
)

func TestSetChatAdministratorCustomTitle(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	mg := (bot.SetChatAdministratorCustomTitle()).SetUserId(Tests{}.Defaults().UserId).SetChatId(1234).SetCustomTitle("Manager")
	messages, err := bot.Send(mg)
	if err != nil {
		fmt.Println(err)
		return
	}
	pretty.Println(messages, string(messages.Result))
	/*m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)*/
	// fmt.Println(*m.Message.MediaGroup)

	/*	message = (bot.MediaGroup()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetMediaGroupId("DQACAgQAAxkDAAECEyJeYaMzAbcxbX80TGvh_mC4AdKCawACNwYAAiWdEVODQvIVVZoJJRgE")
		m, err = bot.Send(message)
		if err != nil {
			fmt.Println(err)
			return
		}*/
	// fmt.Println(*m.Message.MediaGroup)

	/*m, err = (bot.Audio()).SetChatId(Tests{}.Defaults().UserId).SetCaption("Love").SetAudioThumbFilePath("photo.png").SetAudioId("CQACAgQAAxkDAAECEu5eYPQO1_MxAT-J8petsO2WrRRiTAACpwUAAiWdCVO8yR5RNMmjdBgE").Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.Message.Audio)*/
}
