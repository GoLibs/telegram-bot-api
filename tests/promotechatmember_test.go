package tests

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	go_telegram_bot_api "github.com/GoLibs/telegram-bot-api"
)

func TestPromoteChatMember(t *testing.T) {
	bot, err := go_telegram_bot_api.NewTelegramBot("796493295:AAE3EGLAnba_XAsp_ts3sbPTHpW3nitBc4s")
	if err != nil {
		fmt.Println(err)
		return
	}
	up := (bot.PromoteChatMember()).SetUserId(207260097).SetChatId(Tests{}.Defaults().ChatId).EnablePinMessages()
	resp, err := bot.Send(up)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Bool != nil && *resp.Bool {
		fmt.Println("user can now pin messages")
		return
	}
	pretty.Println(resp.Ok, resp.Message)
}
