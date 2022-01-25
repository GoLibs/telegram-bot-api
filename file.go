package tgbotapi

import "io"

type fileInfo struct {
	Field  string
	Path   string
	Reader io.Reader
	Name   string
}
