package tests

// func TestGetUpdates(t *testing.T) {
// 	bot, err := tgbotapi.NewTelegramBot("796493295:AAE3EGLAnba_XAsp_ts3sbPTHpW3nitBc4s")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	updates, err := bot.GetUpdatesChannel(nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	for update := range updates {
// 		m := (bot.Message()).SetChatId(update.Message.Chat.Id).SetText("hello")
// 		bot.Send(m)
//
// 		if update.Message.Video != nil {
// 			pretty.Println(update.Message.Video)
// 			file := (bot.File()).SetFileId(update.Message.Video.FileId)
// 			res, err := bot.Send(file)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			pretty.Println(res)
// 			return
// 		}
// 		bot.StopReceivingUpdates()
// 	}
// 	fmt.Println("g")
// }
