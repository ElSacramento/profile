package middleware

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type loggerKey string

var keyName loggerKey = "logger"

func LoggerFromContext(ctx context.Context) (context.Context, *logrus.Entry) {
	found, ok := ctx.Value(keyName).(*logrus.Entry)
	if ok {
		return ctx, found
	}

	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Logger.SetOutput(os.Stdout)

	ctxWithLogger := context.WithValue(ctx, keyName, logger)
	return ctxWithLogger, logger
}
