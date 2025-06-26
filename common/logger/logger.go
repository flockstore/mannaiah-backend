package logger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a zap.SugaredLogger with a given level.
// If writer is nil, logs go to the default destination.
func New(level string, writer io.Writer) *zap.SugaredLogger {
	parsedLevel := ParseLevel(level)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var core zapcore.Core
	if writer != nil {
		ws := zapcore.AddSync(writer)
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			ws,
			parsedLevel,
		)
	} else {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(parsedLevel)
		cfg.EncoderConfig = encoderCfg

		logger, err := cfg.Build()
		if err != nil {
			panic("cannot initialize logger: " + err.Error())
		}
		return logger.Sugar()
	}

	return zap.New(core).Sugar()
}
