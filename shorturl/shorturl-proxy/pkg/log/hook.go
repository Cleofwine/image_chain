package log

import (
	nativeLog "log"

	"github.com/sirupsen/logrus"
)

type errorHook struct {
}

func (*errorHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (*errorHook) Fire(entry *logrus.Entry) error {
	nativeLog.Println(entry.Message, entry.Data)
	return nil
}
