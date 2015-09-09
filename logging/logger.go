package logging

import (
	"log"
	"os"

	"golang.org/x/net/context"
)

type Logger struct {
	infoLog      *log.Logger
	errorLog     *log.Logger
	InfoLogChan  chan string
	ErrorLogChan chan error
}

// NewLogger returns a new Logger.
func NewLogger() Logger {

	flags := log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lshortfile

	return Logger{
		infoLog:      log.New(os.Stdout, "[INFO]  ", flags),
		errorLog:     log.New(os.Stderr, "[ERROR] ", flags),
		InfoLogChan:  make(chan string),
		ErrorLogChan: make(chan error),
	}

}

// LogInfo sends a message to the InfoLogChan.
func LogInfo(ctx context.Context, info string) {
	FromContext(ctx).InfoLogChan <- info
}

// LogInfo sends an error to the ErrorLogChan.
func LogError(ctx context.Context, err error) {
	FromContext(ctx).ErrorLogChan <- err
}

// Run is used to run the logger. Likely needs to be called as a go routine.
func (l *Logger) Run() {

	for {
		select {
		case info := <-l.InfoLogChan:
			l.infoLog.Println(info)
		case err := <-l.ErrorLogChan:
			l.errorLog.Println(err)
		}
	}

}
