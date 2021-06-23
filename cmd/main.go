package cmd

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/profile/configuration"
	"github.com/profile/service"
)

func main() {
	cfgPath := flag.String("cfg", "", "cfg file path")
	flag.Parse()

	if *cfgPath == "" {
		logrus.Fatalln("Config for service is not set")
	}

	cfg, err := configuration.ParseConfig(cfgPath)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to parse config")
	}

	if err := cfg.ValidateConfig(); err != nil {
		logrus.WithError(err).Fatalln("Failed to validate config")
	}

	srv, err := service.New(cfg)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to initialize service")
	}

	if err := srv.Start(); err != nil {
		logrus.WithError(err).Fatalln("Failed to start service")
	}

	defer func() {
		deadline := time.Now().Add(time.Second * cfg.StopTimeout)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		defer cancel()

		err := srv.Stop(ctx)
		if err == context.DeadlineExceeded {
			err = srv.ForceStop()
		}

		if err != nil {
			logrus.WithError(err).Fatalln("Failed to stop service")
		}
	}()

	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)
	s := <-gracefulStop
	logrus.Printf("Got signal: %v", s)
	close(gracefulStop)
}
