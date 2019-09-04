package log

import (
	"github.com/sirupsen/logrus"
	"os"
)




func New() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}
