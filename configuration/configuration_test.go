package configuration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCfg_ValidateConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cfg := Cfg{
			API:         API{Listen: "localhost:8080"},
			DB:          DB{URL: "postgres"},
			StopTimeout: time.Second,
			Migrations:  Migrations{Directory: "dir"},
		}
		require.Nil(t, cfg.ValidateConfig())
	})
	t.Run("require fields", func(t *testing.T) {
		cfg := Cfg{
			API:        API{},
			Migrations: Migrations{Directory: "dir"},
		}
		err := cfg.ValidateConfig()
		require.NotNil(t, err)
		assert.Contains(t, err.Error(), "API.Listen")
		assert.Contains(t, err.Error(), "StopTimeout")
		assert.Contains(t, err.Error(), "DB.URL")
	})
}
