package driver

import "github.com/sky0621/wolf-bff/src/system"

type Web interface {
}

type web struct {
	cfg system.Config
}

func NewWeb(cfg system.Config) Web {
	return &web{cfg: cfg}
}
