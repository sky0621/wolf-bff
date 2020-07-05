package driver

import "github.com/sky0621/wolf-bff/src/system"

type Web interface {
}

type web struct {
	cfg system.Config
	log system.Logger
}

func NewWeb(cfg system.Config, log system.Logger) Web {
	return &web{cfg: cfg, log: log}
}
