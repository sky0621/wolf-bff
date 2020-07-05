package system

import "github.com/kelseyhightower/envconfig"

type Config interface {
	IsLocal() bool
	GetRDBSetting() RDBSetting
	GetWebSetting() WebSetting
}

type config struct {
	// 起動環境切り分け用
	Env string `default:"local"`

	// DB接続設定用
	RDBHost     string `split_words:"true" default:"localhost"`
	RDBPort     string `split_words:"true" default:"25251"`
	RDBName     string `split_words:"true" default:"wht"`
	RDBUser     string `split_words:"true" default:"postgres"`
	RDBPassword string `split_words:"true" default:"whtpass"`
	RDBSSLMode  string `envconfig:"rdb_sslmode" split_words:"true" default:"disable"`

	// Webサーバ設定用
	WebPort string `default:"8765"`
}

func NewConfig() Config {
	var c config
	if err := envconfig.Process("WHT", &c); err != nil {
		return nil
	}
	return &c
}

func (c *config) IsLocal() bool {
	return c.Env == "local"
}

type RDBSetting struct {
	RDBHost     string
	RDBPort     string
	RDBName     string
	RDBUser     string
	RDBPassword string
	RDBSSLMode  string
}

func (c *config) GetRDBSetting() RDBSetting {
	return RDBSetting{
		RDBHost:     c.RDBHost,
		RDBPort:     c.RDBPort,
		RDBName:     c.RDBName,
		RDBUser:     c.RDBUser,
		RDBPassword: c.RDBPassword,
		RDBSSLMode:  c.RDBSSLMode,
	}
}

type WebSetting struct {
	WebPort string
}

func (c *config) GetWebSetting() WebSetting {
	return WebSetting{
		WebPort: c.WebPort,
	}
}
