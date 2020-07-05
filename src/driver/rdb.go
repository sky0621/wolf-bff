package driver

import (
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/boil"

	"golang.org/x/xerrors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sky0621/wolf-bff/src/system"
)

type RDB interface {
	GetDBWrapper() *sqlx.DB
	Close() error
}

type rdb struct {
	cfg system.Config
	log system.Logger

	dbWrapper *sqlx.DB
}

func NewRDB(cfg system.Config, log system.Logger) (RDB, error) {
	stg := cfg.GetRDBSetting()

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		stg.RDBHost, stg.RDBPort, stg.RDBName, stg.RDBUser, stg.RDBPassword, stg.RDBSSLMode)

	dbWrapper, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	boil.SetLocation(loc)

	return &rdb{cfg: cfg, log: log, dbWrapper: dbWrapper}, nil
}

func (r *rdb) GetDBWrapper() *sqlx.DB {
	return r.dbWrapper
}

func (r *rdb) Close() error {
	if r == nil {
		return nil
	}
	if r.dbWrapper == nil {
		return nil
	}
	return r.dbWrapper.Close()
}
