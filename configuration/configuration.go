package configuration

import (
	"time"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type API struct {
	Listen string `yaml:"listen" validate:"required"`
}

type DB struct {
	URL string `yaml:"url" validate:"required"`
}

type Migrations struct {
	Directory string `yaml:"directory" validate:"required"`
}

type Cfg struct {
	API         API           `yaml:"api" validate:"required"`
	DB          DB            `yaml:"db" validate:"required"`
	StopTimeout time.Duration `yaml:"stop_timeout" validate:"required"`
	Migrations  Migrations    `yaml:"migrations" validate:"required"`
}

// todo: think about smth more useful
func New(vp *viper.Viper) Cfg {
	return Cfg{
		API: API{
			Listen: vp.GetString("api_listen"),
		},
		DB: DB{
			URL: vp.GetString("db_url"),
		},
		StopTimeout: vp.GetDuration("stop_timeout"),
		Migrations: Migrations{
			Directory: vp.GetString("migrations_directory"),
		},
	}
}

func (c *Cfg) ValidateConfig() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}
