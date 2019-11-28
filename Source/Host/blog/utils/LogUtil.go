package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

const (
	YEAR  = "Y"
	MONTH = "M"
	DAY   = "D"
)

const (
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
)

type Logger struct {
	MailSender         func(title, content string, to, cc, bcc []string)
	Reminder           []string
	InfoPathFormatter  string
	WarnPathFormatter  string
	ErrorPathFormatter string
	Level              string
	Debug              bool
	SplitType          string
	SystemName         string
	errorLogger        *log.Logger
	warnLogger         *log.Logger
	infoLogger         *log.Logger
	lastUpdateTimeFlag int
}

// region get logrus's logger 获取logrus的logger

/**
get logrus's logger
根据配置项获取logrus的logger
*/
func getLog(debug bool, filePathFormatter string, level log.Level) *log.Logger {
	logger := log.New()
	logger.SetLevel(level)
	if !debug {
		logFile := getFile(filePathFormatter)
		logger.SetOutput(logFile)
		logger.SetFormatter(&log.TextFormatter{
			ForceColors:               false,
			DisableColors:             false,
			EnvironmentOverrideColors: false,
			DisableTimestamp:          false,
			FullTimestamp:             true,
			TimestampFormat:           "2006/01/02 15:04:05",
			DisableSorting:            false,
			SortingFunc:               nil,
			DisableLevelTruncation:    false,
			QuoteEmptyFields:          false,
			FieldMap:                  nil,
		})
	} else {
		logger.SetOutput(os.Stdout)
		logger.SetFormatter(&log.TextFormatter{
			ForceColors:               true,
			DisableColors:             false,
			EnvironmentOverrideColors: true,
			DisableTimestamp:          false,
			FullTimestamp:             true,
			TimestampFormat:           "2006/01/02 15:04:05",
			DisableSorting:            false,
			SortingFunc:               nil,
			DisableLevelTruncation:    false,
			QuoteEmptyFields:          false,
			FieldMap:                  nil,
		})
	}
	return logger
}

/**
Create log directory and create log file
创建日志文件夹和日志文件
*/
func getFile(filePath string) *os.File {
	if err := os.MkdirAll(path.Dir(filePath), 0755); err != nil {
		if !os.IsExist(err) {
			log.Fatalln("error create dir")
		}
	}
	logFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(filePath)
		log.Fatalln("error open errorLog file")
	}
	return logFile
}

/**
format log file name
计算文件名
*/
func formatLogFileName(format string, now time.Time) string {
	year := fmt.Sprintf("%000d", now.Year())
	month := fmt.Sprintf("%00d", int(now.Month()))
	day := fmt.Sprintf("%00d", now.Day())
	hour := fmt.Sprintf("%00d", now.Hour())
	min := fmt.Sprintf("%00d", now.Minute())
	fileName := strings.Replace(format, "%Y", year, -1)
	fileName = strings.Replace(fileName, "%M", month, -1)
	fileName = strings.Replace(fileName, "%D", day, -1)
	fileName = strings.Replace(fileName, "%H", hour, -1)
	fileName = strings.Replace(fileName, "%m", min, -1)
	return fileName
}

func getTimeUpdateFlag(key string, now time.Time) int {
	//用switch判断比用hash表驱动更快,编译的时候会自动生成跳转表
	switch key {
	case "Y":
		return now.Year()
	case "M":
		return int(now.Month())
	case "D":
		return now.Day()
	case "H":
		return now.Hour()
	case "m":
		return now.Minute()
	default:
		return now.Day()
	}
}

// endregion

func (l *Logger) InitLogger() {
	if l.InfoPathFormatter == "" {
		l.InfoPathFormatter = "./accesslog/info-YYYY-MM-DD.log"
	}
	if l.WarnPathFormatter == "" {
		l.WarnPathFormatter = "./accesslog/warn-YYYY-MM-DD.log"
	}
	if l.ErrorPathFormatter == "" {
		l.ErrorPathFormatter = "./accesslog/error-YYYY-MM-DD.log"
	}
	if l.SplitType == "" {
		l.SplitType = "D"
	}
	if l.SystemName == "" {
		l.SystemName = "BugTian's Blog"
	}
	if l.Level == "" {
		l.Level = INFO
	}
	now := time.Now()
	levelMap := map[string]log.Level{
		"info":  log.InfoLevel,
		"warn":  log.WarnLevel,
		"error": log.ErrorLevel,
	}
	l.errorLogger = getLog(l.Debug, formatLogFileName(l.ErrorPathFormatter, now), levelMap[l.Level])
	l.warnLogger = getLog(l.Debug, formatLogFileName(l.WarnPathFormatter, now), levelMap[l.Level])
	l.infoLogger = getLog(l.Debug, formatLogFileName(l.InfoPathFormatter, now), levelMap[l.Level])
	l.lastUpdateTimeFlag = getTimeUpdateFlag(l.SplitType, now)
}

func (l *Logger) checkPath() {
	now := time.Now()
	timeFlag := getTimeUpdateFlag(l.SplitType, now)
	if timeFlag != l.lastUpdateTimeFlag {
		l.infoLogger.SetOutput(getFile(formatLogFileName(l.InfoPathFormatter, now)))
		l.warnLogger.SetOutput(getFile(formatLogFileName(l.WarnPathFormatter, now)))
		l.errorLogger.SetOutput(getFile(formatLogFileName(l.ErrorPathFormatter, now)))
		l.lastUpdateTimeFlag = timeFlag
	}
}

func (l *Logger) LogInfo(format string, a ...interface{}) {
	l.checkPath()
	l.infoLogger.Infof(format, a)
}

func (l *Logger) LogWarn(format string, a ...interface{}) {
	l.checkPath()
	l.warnLogger.Warnf(format, a)
}

func (l *Logger) LogError(format string, a ...interface{}) {
	l.checkPath()
	l.errorLogger.Errorf(format, a)
	if !l.Debug && l.MailSender != nil && l.Reminder != nil && len(l.Reminder) > 0 {
		l.MailSender("System:["+l.SystemName+"] Error", fmt.Sprintf(format, a), l.Reminder, nil, nil)
	}
}
