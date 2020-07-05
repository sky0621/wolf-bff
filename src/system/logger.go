package system

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
}

type logger struct {
	l zerolog.Logger
}

func NewLogger(cfg Config) Logger {
	l := zerolog.New(os.Stdout)
	return &logger{l: l}
}
