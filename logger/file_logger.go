package logger

import (
	"log"
)

type FileLogger struct {
}

// method
func (fl *FileLogger) Log(msg string) {
	log.SetPrefix("[File] ")
	log.Println(msg)
}

// function
func NewFileLogger() *FileLogger {
	f := &FileLogger{}
	return f
}


/*
example for how this work flow in other language
class FileLogger implements Logger {

	@override
	void Log(msg string) {

	}
}
*/
