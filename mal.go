package mal

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

type Logger struct{}

var logger = &Logger{}

var (
	infoColor  = color.New(color.FgBlue).SprintFunc()
	warnColor  = color.New(color.FgYellow).SprintFunc()
	errorColor = color.New(color.FgRed).SprintFunc()
	debugColor = color.New(color.FgGreen).SprintFunc()
)

type logLevel struct {
	prefix string
	color  func(a ...interface{}) string
}

var (
	infoLevel  = logLevel{prefix: "[INFO]", color: infoColor}
	warnLevel  = logLevel{prefix: "[WARN]", color: warnColor}
	errorLevel = logLevel{prefix: "[ERROR]", color: errorColor}
	debugLevel = logLevel{prefix: "[DEBUG]", color: debugColor}
)

func (l *Logger) log(level logLevel, file string, line int, args ...interface{}) {
	filePath := strings.Split(file, "/")
	fileName := fmt.Sprintf("%s/%s:%d", filePath[len(filePath)-2], filePath[len(filePath)-1], line)
	link := fmt.Sprintf("file://%s:%d", file, line)
	linkText := fileName
	ansiLink := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", link, linkText)

	fmt.Printf("%s %s %s\n", level.color(level.prefix), ansiLink, args)
}

func log(level logLevel, args ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}
	logger.log(level, file, line, args)
}

func Info(args ...interface{}) {
	log(infoLevel, args)
}

func Warn(args ...interface{}) {
	log(warnLevel, args)
}

func Error(args ...interface{}) {
	log(errorLevel, args)
}

func Debug(args ...interface{}) {
	log(debugLevel, args)
}

// :)
