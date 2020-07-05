package driver

import "github.com/sky0621/wolf-bff/src/system"

type RDB interface {
}

type rdb struct {
	cfg system.Config
}

func NewRDB(cfg system.Config) RDB {
	return &rdb{cfg: cfg}
}
