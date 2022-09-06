package log

import "github.com/sirupsen/logrus"

var L *logrus.Logger

func init() {
	L = logrus.New()
	L.SetLevel(logrus.DebugLevel)
	L.SetFormatter(&logrus.TextFormatter{})
}
