package log

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Setup() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	if viper.GetString("APP_LOG_LEVEL") == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}

}
