package log

import (
	"gtk-lw-server-main/util"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// Set the log format
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	// Create a file log
	file, err := os.OpenFile(util.PathLog+"gtk-lw.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Info("Failed to log to file, using default stderr")
	} else {
		logrus.SetOutput(file)
	}

}

// Info logs a message at the Info level.
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// InfoWithFields logs a message at the Info level with additional fields.
func PacketSent(deviceid string, format string, buffer []byte) {
	logrus.WithFields(logrus.Fields{"imei": deviceid}).Infof("packet sent: "+format, buffer)
}

func PacketReceived(deviceid string, format string, buffer []byte) {
	logrus.WithFields(logrus.Fields{"imei": deviceid}).Infof("packet received: "+format, buffer)
}

// Error logs a message at the Error level.
func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	logrus.Error(args...)
}
