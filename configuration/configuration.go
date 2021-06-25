package configuration

import (
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type API struct {
	Listen string `validate:"required"`
}

type DB struct {
	URL string `validate:"required"`
}

type Migrations struct {
	Directory string `validate:"required"`
}

type Notify []string

type Cfg struct {
	API         API           `validate:"required"`
	DB          DB            `validate:"required"`
	StopTimeout time.Duration `validate:"required"`
	Migrations  Migrations    `validate:"required"`
	Notify      Notify
}

func New(vp *viper.Viper) Cfg {
	// get everything from viper and use default values
	// notice: unmarshal didn't work, but maybe the way of use was wrong
	subscribers := make([]string, 0)
	notify := vp.GetString("notify")
	if notify != "" {
		notify = strings.TrimLeft(notify, "[")
		notify = strings.TrimRight(notify, "]")
		subscribers = strings.Split(notify, ",")
	}

	timeout := vp.GetDuration("stop_timeout")
	if timeout == 0 {
		timeout = time.Second * 30
	}

	migrationsDir := vp.GetString("migrations_directory")
	if migrationsDir == "" {
		migrationsDir = "/etc/migrations"
	}

	cfg := Cfg{
		API: API{
			Listen: vp.GetString("api_listen"),
		},
		DB: DB{
			URL: vp.GetString("db_url"),
		},
		StopTimeout: timeout,
		Migrations: Migrations{
			Directory: migrationsDir,
		},
		Notify: subscribers,
	}
	return cfg
}

func (c *Cfg) ValidateConfig() error {
	// todo: logical validation for values
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}
