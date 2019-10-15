package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"sync"
)

var logger *logrus.Logger
var once sync.Once

// Create Singleton instance of the logger
func Instance() *logrus.Logger {
	once.Do(func() {
		logger = &logrus.Logger{
			Out:   os.Stderr,
			Level: logrus.DebugLevel,
			Formatter: &easy.Formatter{
				TimestampFormat: "2006-01-02 15:04:05",
				LogFormat:       "[%lvl%]: %time% - %msg%",
			},
		}
	})
	return logger
}
