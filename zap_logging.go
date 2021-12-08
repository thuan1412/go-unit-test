package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price uint
}

func hello2(name string) string {
	if name == "" {
		return "Hello, what's your name?"
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}

func welcome2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Hello world")
}

func New2(a string) string {
	return a
}

var logger *zap.Logger

// func main() {
// 	InitLogger()
// }

type foo struct {
	One string
	Two string
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger.Error("Logger initialized")
	logger.Debug("test", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Info("test", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Warn("test", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Info("gorm data: product", zap.Any("ads", Product{}))
	logger.Error("test", zap.Any("foo", foo{One: "one", Two: "two"}), zap.String("asd", "123"))

	logger.Info("Success..",
		zap.String("statusCode", "200"))
}

