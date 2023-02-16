package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type Response struct {
	Ok                bool                       `json:"ok"`
	ErrorCode         int64                      `json:"error_code,omitempty"`
	Description       string                     `json:"description,omitempty"`
	Result            json.RawMessage            `json:"result"`
	Bool              *bool                      `json:"-"`
	Int               *int64                     `json:"-"`
	Message           *structs.Message           `json:"-"`
	Messages          []structs.Message          `json:"-"`
	UserProfilePhotos *structs.UserProfilePhotos `json:"-"`
	Updates           []Update                   `json:"-"`
	ChatMembers       []structs.ChatMember       `json:"-"`
	ChatMember        *structs.ChatMember        `json:"-"`
	Chat              *structs.Chat              `json:"-"`
	User              *structs.User              `json:"-"`
	File              *structs.File              `json:"-"`
	StickerSet        *structs.StickerSet        `json:"-"`
	MessageId         *structs.MessageId         `json:"-"`
	Raw               []byte                     `json:"-"`
}
