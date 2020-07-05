package system

import "github.com/rs/zerolog"

type Logger interface {
}

type logger struct {
	l *zerolog.Logger
}

func NewLogger(l *zerolog.Logger) Logger {
	return &logger{l: l}
}
