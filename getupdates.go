package go_telegram_bot_api

import (
	"encoding/json"

	"github.com/golibs/telegram-bot-api/responses"

	"github.com/golibs/telegram-bot-api/structs"
)

type getUpdates struct {
	parent         *TelegramBot
	buffer         int64
	offset         int64
	limit          int64
	timeout        int64
	allowedUpdates []string
}

func (gu *getUpdates) Offset() int64 {
	return gu.offset
}

func (gu *getUpdates) Limit() int64 {
	return gu.limit
}

func (gu *getUpdates) Timeout() int64 {
	return gu.timeout
}

func (gu *getUpdates) AllowedUpdates() []string {
	return gu.allowedUpdates
}

func (gu *getUpdates) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Offset         int64    `json:"offset"`
		Limit          int64    `json:"limit"`
		Timeout        int64    `json:"timeout"`
		AllowedUpdates []string `json:"allowed_updates"`
	}{
		Offset:         gu.offset,
		Limit:          gu.limit,
		Timeout:        gu.timeout,
		AllowedUpdates: gu.allowedUpdates,
	})
}

func (gu *getUpdates) response() interface{} {
	return &[]structs.Update{}
}

func (gu *getUpdates) method() string {
	return "POST"
}

func (gu *getUpdates) endpoint() string {
	return "getUpdates"
}

func (gu *getUpdates) SetBuffer(buffer int64) *getUpdates {
	gu.buffer = buffer
	return gu
}

func (gu *getUpdates) Set(updates *getUpdates) *getUpdates {
	gu.offset = updates.offset
	gu.limit = updates.limit
	gu.allowedUpdates = updates.allowedUpdates
	gu.timeout = updates.timeout
	return gu
}

func (gu *getUpdates) SetOffset(offset int64) *getUpdates {
	gu.offset = offset
	return gu
}

func (gu *getUpdates) SetLimit(limit int64) *getUpdates {
	gu.limit = limit
	return gu
}

func (gu *getUpdates) SetTimeout(timeout int64) *getUpdates {
	gu.timeout = timeout
	return gu
}

func (gu *getUpdates) SetAllowedUpdates(allowed []string) *getUpdates {
	gu.allowedUpdates = allowed
	return gu
}

func (gu *getUpdates) Get() (result *responses.Response, err error) {
	result, err = gu.parent.Send(gu)
	return
}
