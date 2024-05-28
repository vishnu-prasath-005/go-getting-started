package logger

import (
	"log"
)

type FileLogger struct {
}

//method
func (fl *FileLogger) Log(msg string) {
	log.SetPrefix("[File] ")
	log.Println(msg)
}

//function
func NewFileLogger() *FileLogger {
	return &FileLogger{}
}
