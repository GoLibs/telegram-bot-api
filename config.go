package tgbotapi

type Config interface {
	method() string
	endpoint() string
	response() interface{}
	marshalJSON() ([]byte, error)
}

type file interface {
	medias() []fileInfo
}
