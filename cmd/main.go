package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/profile/configuration"
	"github.com/profile/middleware"
	"github.com/profile/service"
)

// todo: make function smaller
func main() {
	ctx, logger := middleware.LoggerFromContext(context.Background())

	vp := viper.NewWithOptions(viper.KeyDelimiter("_"))
	vp.SetEnvPrefix("profile")
	vp.AutomaticEnv()

	cfg := configuration.New(vp)

	if err := cfg.ValidateConfig(); err != nil {
		logger.WithError(err).Fatalln("Failed to validate config")
	}

	srv, err := service.New(ctx, cfg)
	if err != nil {
		logger.WithError(err).Fatalln("Failed to initialize service")
	}

	if err := srv.Start(); err != nil {
		logger.WithError(err).Fatalln("Failed to start service")
	}

	defer func() {
		deadline := time.Now().Add(time.Second * cfg.StopTimeout)
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()

		err := srv.Stop(ctx)
		if err == context.DeadlineExceeded {
			err = srv.ForceStop()
		}

		if err != nil {
			logger.WithError(err).Fatalln("Failed to stop service")
		}
	}()

	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)
	s := <-gracefulStop
	logger.Infof("Got signal: %v", s)
	close(gracefulStop)
}
