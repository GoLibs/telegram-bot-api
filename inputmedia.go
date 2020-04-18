package go_telegram_bot_api

type inputMedia interface {
	medias() []fileInfo
	marshalJSON() ([]byte, error)
}
