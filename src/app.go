package wht

import (
	"github.com/sky0621/wolf-bff/src/driver"
	"github.com/sky0621/wolf-bff/src/system"
	"golang.org/x/xerrors"
)

type App interface {
	Start() error
	Shutdown() error
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
	if a == nil {
		return xerrors.New("app is nil")
	}
	return a.web.Start()
}

func (a *app) Shutdown() error {
	if a == nil {
		return xerrors.New("app is nil")
	}
	return a.rdb.Close()
}
