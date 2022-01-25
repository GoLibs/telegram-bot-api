package tests

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestKickChatMember(t *testing.T) {
	bot, err := tgbotapi.NewTelegramBot("796493295:AAE3EGLAnba_XAsp_ts3sbPTHpW3nitBc4s")
	if err != nil {
		fmt.Println(err)
		return
	}
	up := (bot.KickChatMember()).SetUserId(207260097).SetChatId(Tests{}.Defaults().ChatId)
	resp, err := bot.Send(up)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Bool != nil && *resp.Bool {
		fmt.Println("kicked")
		return
	}
	pretty.Println(resp.Ok, resp.Message)
}
