package main

import (
	"fmt"

	"github.com/vishu-prasath-005/go-getting-started/logger"
)

var clog logger.Logger = logger.NewStdLogger()
var flog logger.Logger = logger.NewFileLogger()

func main() {
	i := 0
	clog.Log("Hai this in console logger")
	flog.Log("Hai this is in the file logger")

	// this defer executes last
	defer func(initial int) {
		fmt.Println("In teh defer function", initial)
	}(i)

	i++
	// this defer executes first
	defer func(initial int) {
		fmt.Println("In teh defer function", initial)
	}(i)

	fmt.Println("In teh outside defer function", i)

}
