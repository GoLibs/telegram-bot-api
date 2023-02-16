package tgbotapi

import "time"

type SendOptions struct {
	timeout time.Duration
}

func NewSendOptions() *SendOptions {
	return &SendOptions{}
}

func (s *SendOptions) SetTimeout(dur time.Duration) *SendOptions {
	s.timeout = dur

	return s
}
