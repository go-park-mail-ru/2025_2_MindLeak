package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

var Log *logrus.Logger

type contextKey string

const RequestIDKey contextKey = "requestID"

func Init() *logrus.Logger {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&CustomFormatter{})

	return Log
}

func logWithContext(ctx context.Context) *logrus.Entry {
	reqID, _ := ctx.Value(RequestIDKey).(string)
	if reqID == "" {
		reqID = "unknown"
	}

	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return Log.WithField("requestID", reqID)
	}
	fn := runtime.FuncForPC(pc)
	funcPath := fn.Name()
	lastSlash := strings.LastIndex(funcPath, "/")
	if lastSlash == -1 {
		lastSlash = 0
	}
	parts := strings.Split(funcPath[lastSlash:], ".")
	pkg := "unknown"
	funcName := "anonymous"
	if len(parts) >= 2 {
		pkg = strings.TrimLeft(parts[0], "/")
		funcName = strings.Join(parts[1:], ".")
	}

	return Log.WithFields(logrus.Fields{
		"requestID": reqID,
		"package":   pkg,
		"function":  funcName,
	})
}

func Info(ctx context.Context, format string, args ...interface{}) {
	logWithContext(ctx).Infof(format, args...)
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	logWithContext(ctx).Warnf(format, args...)
}

func Error(ctx context.Context, format string, args ...interface{}) {
	logWithContext(ctx).Errorf(format, args...)
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	logWithContext(ctx).Debugf(format, args...)
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var color string
	switch entry.Level {
	case logrus.InfoLevel:
		color = "\033[34m" // Синий
	case logrus.WarnLevel:
		color = "\033[33m" // Жёлтый
	case logrus.ErrorLevel:
		color = "\033[31m" // Красный
	case logrus.FatalLevel, logrus.PanicLevel:
		color = "\033[35m" // Фиолетовый
	default:
		color = "\033[0m"
	}

	level := fmt.Sprintf("%s[%s]\033[0m", color, strings.ToUpper(entry.Level.String()))
	loc, _ := time.LoadLocation("Europe/Moscow")
	ts := entry.Time.In(loc).Format("2006-01-02 15:04:05")

	reqID := entry.Data["requestID"]
	pkg := entry.Data["package"]
	fn := entry.Data["function"]

	message := fmt.Sprintf("%s[%s][%v][%v][%v] %s\n",
		level,
		ts,
		reqID,
		pkg,
		fn,
		entry.Message,
	)

	return []byte(message), nil
}
