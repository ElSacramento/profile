package configuration

import (
	"time"

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

type NotifyAddr struct {
	Addr string
}

type Notify []NotifyAddr

type Cfg struct {
	API         API           `validate:"required"`
	DB          DB            `validate:"required"`
	StopTimeout time.Duration `json:"stop_timeout" validate:"required"`
	Migrations  Migrations    `validate:"required"`
	Notify      Notify
}

func (c *Cfg) ValidateConfig() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}
