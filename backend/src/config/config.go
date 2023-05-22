package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	ServerAddress string
	ServerPort    string
	DbConfig      DbConfig `mapstructure:"db"`
	Logger        *logrus.Logger
}

type DbConfig struct {
    ServerUrl     string
}

func LoadConfigFromEnv() (Config, error) {
	var cfg Config

	cfg.ServerAddress = "localhost"
	cfg.ServerPort = "8599"

	cfg.DbConfig.ServerUrl= os.Getenv("SERVER_URL")

	logger := getLogger("loglevel")

	cfg.Logger = logger
	return cfg, nil
}

func getLogger(loglevel string) *logrus.Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	level, ok := logrus.ParseLevel(loglevel)
	if ok != nil {
		level = logrus.DebugLevel
	}
	logrusLogger.SetLevel(level)
	return logrusLogger
}
