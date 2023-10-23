package config

import (
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
)


func NewLogger(logFileName string) *zap.Logger {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   LogRootPath() +"/"+logFileName,
		MaxSize:    LogMaxFileSize(),
		MaxBackups: LogMaxBackups(),
		MaxAge:     LogMaxAge(),
		Compress:   true,
	}
	// Create a zap core that writes logs to the lumberjack logger
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(lumberjackLogger),
		zap.DebugLevel,
	)
	// Create a logger with the zap core
	logger := zap.New(zapCore)
	return logger
}

func LogRequestPath(ctx *fiber.Ctx) zap.Field {
	return zap.String("path", ctx.Path())
}

func LogRootPath() string {
	LoadDotEnv()
	return os.Getenv("LOG_ROOT_PATH")
}

func LogMaxFileSize() int {
	LoadDotEnv()
	return helpers.ConvertToInt(os.Getenv("LOG_MAX_SIZE"))
}

func LogMaxAge() int {
	LoadDotEnv()
	return helpers.ConvertToInt(os.Getenv("LOG_MAX_AGE"))
}

func LogMaxBackups() int {
	LoadDotEnv()
	return helpers.ConvertToInt(os.Getenv("LOG_MAX_BACKUPS"))
}


