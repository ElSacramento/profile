package configuration

import (
	"io/ioutil"
	"time"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
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

func ParseConfig(cfgPath *string) (Cfg, error) {
	data, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		return Cfg{}, err
	}
	var cfg Cfg
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Cfg{}, err
	}
	return cfg, nil
}

func (c *Cfg) ValidateConfig() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}
