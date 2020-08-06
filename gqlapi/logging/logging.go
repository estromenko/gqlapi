package logging

import (
	"gqlapi/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLevel(lvl string) zap.Option {
	levels := map[string]zapcore.Level{
		"info":  zapcore.InfoLevel,
		"debug": zapcore.DebugLevel,
		"error": zapcore.ErrorLevel,
		"fatal": zapcore.FatalLevel,
	}
	return zap.IncreaseLevel(levels[lvl])
}

// NewLogger ...
func NewLogger(conf *config.Config) (*zap.Logger, error) {
	level := getLevel(conf.Logger.Level)
	logger, err := zap.NewDevelopment(level)

	logger.Info("Logger configured successfully")
	return logger, err
}
