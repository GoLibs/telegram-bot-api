package go_telegram_bot_api

import "io"

type fileInfo struct {
	Field  string
	Path   string
	Reader io.Reader
	Name   string
}
