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

func (l *Logger) log(level logLevel, text, file string, line int) {
	filePath := strings.Split(file, "/")
	fileName := fmt.Sprintf("%s/%s:%d", filePath[len(filePath)-2], filePath[len(filePath)-1], line)
	link := fmt.Sprintf("file://%s:%d", file, line)
	linkText := fileName
	ansiLink := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", link, linkText)

	fmt.Printf("%s %s %s\n", level.color(level.prefix), ansiLink, text)
}

func log(level logLevel, text string) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}
	logger.log(level, text, file, line)
}

func Info(text string) {
	log(infoLevel, text)
}

func Warn(text string) {
	log(warnLevel, text)
}

func Error(text string) {
	log(errorLevel, text)
}

func Debug(text string) {
	log(debugLevel, text)
}
