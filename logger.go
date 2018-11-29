package logger

import (
	"fmt"
)

var log LogInterface

func InitLogger(name string,config map[string]string) (err error){
	switch name{
	case "file":
		log,err = NewFileLogger(config)
	case "console":
		log,err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unsupport logger name:%s",name)
	}

	return
}

func Debug(format string,args ...interface{}){
	log.Debug(format,args...)
}

func Trace(format string,args ...interface{}){
	log.Trace(format,args...)
}

func Info(format string,args ...interface{}){
	log.Info(format,args...)
}

func Error(format string,args ...interface{}){
	log.Error(format,args...)
}

func Fatal(format string,args ...interface{}){
	log.Fatal(format,args...)
}

func Warn(format string,args ...interface{}){
	log.Warn(format,args...)
}