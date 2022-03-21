package main

import (
	"time"

	"github.com/Sorarinu/golang-gin-sample/app/v1/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func EncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func newEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "timestamp"
	cfg.LevelKey = "severity"
	cfg.MessageKey = "message"
	cfg.EncodeLevel = EncodeLevel
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}

func NewCloudLoggingLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.InitialFields = map[string]interface{}{"appVersion": "v1.0.0"}
	cfg.EncoderConfig = newEncoderConfig()
	logger, _ := cfg.Build()
	return logger
}

func main() {
	router := gin.Default()

	// CORS 許可
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"*",
	}
	config.AllowMethods = []string{
		"POST",
		"PUT",
		"DELETE",
		"GET",
		"OPTIONS",
	}
	config.AllowHeaders = []string{
		"Content-Type",
		"Content-Length",
		"Authorization",
		"X-CSRF-Token",
		"Accept-Encoding",
		"Access-Control-Allow-Headers",
	}
	config.MaxAge = 24 * time.Hour
	router.Use(cors.New(config))

	logger := NewCloudLoggingLogger()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// zap.L().Info("Info Message")

	v1 := router.Group("/v1")
	{
		ping := controllers.NewPingController()
		v1.GET("ping", ping.Pong)
	}

	router.Run(":8080")
}
