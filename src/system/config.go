package system

import "github.com/kelseyhightower/envconfig"

type Config interface {
	IsLocal() bool
	GetWebSetting() WebSetting
}

type config struct {
	Env     string `default:"local"`
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

type WebSetting struct {
	WebPort string
}

func (c *config) GetWebSetting() WebSetting {
	return WebSetting{
		WebPort: c.WebPort,
	}
}
