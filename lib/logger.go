package lib

import (
	"fmt"
	"log"
	"os"
)

// Logger levels
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var (
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warnLogger  = log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Current log level (default: INFO)
	currentLevel = INFO
)

// SetLogLevel sets the current logging level
func SetLogLevel(level int) {
	currentLevel = level
}

// Debug logs debug messages
func Debug(v ...interface{}) {
	if currentLevel <= DEBUG {
		debugLogger.Output(2, fmt.Sprint(v...))
	}
}

// Debugf logs formatted debug messages
func Debugf(format string, v ...interface{}) {
	if currentLevel <= DEBUG {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Info logs info messages
func Info(v ...interface{}) {
	if currentLevel <= INFO {
		infoLogger.Output(2, fmt.Sprint(v...))
	}
}

// Infof logs formatted info messages
func Infof(format string, v ...interface{}) {
	if currentLevel <= INFO {
		infoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Warn logs warning messages
func Warn(v ...interface{}) {
	if currentLevel <= WARN {
		warnLogger.Output(2, fmt.Sprint(v...))
	}
}

// Warnf logs formatted warning messages
func Warnf(format string, v ...interface{}) {
	if currentLevel <= WARN {
		warnLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Error logs error messages
func Error(v ...interface{}) {
	if currentLevel <= ERROR {
		errorLogger.Output(2, fmt.Sprint(v...))
	}
}

// Errorf logs formatted error messages
func Errorf(format string, v ...interface{}) {
	if currentLevel <= ERROR {
		errorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}
