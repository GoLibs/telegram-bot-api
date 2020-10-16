package tools

import "github.com/GoLibs/telegram-bot-api/structs"

type Keyboards struct {
}

func (k Keyboards) NewInlineKeyboard() *structs.InlineKeyboardMarkup {
	return &structs.InlineKeyboardMarkup{}
}

func (k Keyboards) NewInlineKeyboardFromSlicesOfMaps(slicesOfMaps [][]map[string]string) *structs.InlineKeyboardMarkup {
	return (&structs.InlineKeyboardMarkup{}).FromSlicesOfMaps(slicesOfMaps)
}

func (k Keyboards) NewReplyKeyboardFromSlicesOfStrings(slicesOfStrings [][]string) *structs.ReplyKeyboardMarkup {
	return (&structs.ReplyKeyboardMarkup{}).FromSlicesOfStrings(slicesOfStrings)
}
