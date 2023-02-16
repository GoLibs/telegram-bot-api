package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendMediaGroup(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bot)
	/*mg := (bot.MediaGroup()).SetChatId(Tests{}.Defaults().UserId).AddPhotoId("AgACAgQAAxkDAAECEyteYb4740re3aRWitYXHsqL0_eAUQACO7QxGyWdEVMRRSHBZjoRmXBXqBsABAEAAwIAA3gAA63FCAABGAQ", "first").AddPhotoId("AgACAgQAAxkDAAECEypeYb47n_RhMypebJ7S19g_JAk96AACEbQxGyWdEVM2I7m0B4Fz0WU_qBsABAEAAwIAA3gAAxfJCAABGAQ", "second").AddVideoId("BAACAgQAAxkDAAECEyxeYb47bqJ7Hftk7FXxCkAnDTmVJgACPAYAAiWdEVOQokMVflxOPhgE", "").AddVideoId("BAACAgQAAxkDAAECEyxeYb47bqJ7Hftk7FXxCkAnDTmVJgACPAYAAiWdEVOQokMVflxOPhgE", "")
	messages, err := bot.Send(mg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(messages, string(messages.Result))*/
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
