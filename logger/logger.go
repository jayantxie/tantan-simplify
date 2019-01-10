package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type LoggerGroup struct {
	AccessLogger *logrus.Logger
	ErrorLogger  *logrus.Logger
}

// TODO: 日志命名提起成配置项（matrix 组件的运行时环境也应该统一）
const (
	accessLogName = "access.log"
	errorLogName  = "error.log"
)

var Logger = LoggerGroup{
	AccessLogger: logrus.New(),
	ErrorLogger:  logrus.New(),
}

// Reopen log fd handlers when receiving signal syscall.SIGUSR1
func ReopenLogs(logDir string) error {
	if logDir == "" {
		return nil
	}

	accessLog, err := os.OpenFile(path.Join(logDir, accessLogName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	errorLog, err := os.OpenFile(path.Join(logDir, errorLogName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	oldFd := Logger.AccessLogger.Out.(*os.File)
	Logger.AccessLogger.Out = accessLog
	oldFd.Close()

	oldFd = Logger.ErrorLogger.Out.(*os.File)
	Logger.ErrorLogger.Out = errorLog
	oldFd.Close()

	return nil
}

// @backtrackLevel: log the backtrack info when logging level is >= backtrackLevel
func MustSetup(logDir, logLevel, backtrackLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	btLevel, err := logrus.ParseLevel(backtrackLevel)
	if err != nil {
		return err
	}

	formatter := new(logrus.JSONFormatter)
	Logger.AccessLogger.Formatter = formatter
	Logger.ErrorLogger.Formatter = formatter

	Logger.ErrorLogger.Level = level
	Logger.ErrorLogger.Hooks.Add(NewBackTrackHook(btLevel))

	if logDir == "" {
		Logger.AccessLogger.Out = os.Stdout
		Logger.ErrorLogger.Out = os.Stderr
		return nil
	}
	accessLog, err := os.OpenFile(path.Join(logDir, accessLogName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	errorLog, err := os.OpenFile(path.Join(logDir, errorLogName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	Logger.AccessLogger.Out = accessLog
	Logger.ErrorLogger.Out = errorLog

	return nil
}

func GetHTTPLogger(c *gin.Context) *logrus.Entry {
	reqID := c.GetString("req_id")
	if reqID == "" {
		return logrus.NewEntry(Logger.ErrorLogger)
	}
	return Logger.ErrorLogger.WithField("req_id", reqID)
}
