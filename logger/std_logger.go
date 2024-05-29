package logger

import "log"

type StdLogger struct {
}

// method
func (sl *StdLogger) Log(msg string) {
	log.SetPrefix("[CONSOLE] ")
	log.Println(msg)
}

// function
func NewStdLogger() *StdLogger {
	return &StdLogger{}
}
