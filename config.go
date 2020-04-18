package go_telegram_bot_api

type config interface {
	method() string
	endpoint() string
	response() interface{}
	marshalJSON() ([]byte, error)
}

type file interface {
	medias() []fileInfo
}
