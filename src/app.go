package wht

import (
	"github.com/sky0621/wolf-bff/src/driver"
	"github.com/sky0621/wolf-bff/src/system"
)

type App interface {
	Start() error
	Shutdown()
}

type app struct {
	cfg system.Config
	log system.Logger
	rdb driver.RDB
	web driver.Web
}

func NewApp(cfg system.Config, log system.Logger, rdb driver.RDB, web driver.Web) App {
	return &app{cfg: cfg, log: log, rdb: rdb, web: web}
}

func (a *app) Start() error {
	// FIXME:
	return nil
}

func (a *app) Shutdown() {
	// FIXME:
}
