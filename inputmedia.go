package tgbotapi

type inputMedia interface {
	medias() []fileInfo
	marshalJSON() ([]byte, error)
}
