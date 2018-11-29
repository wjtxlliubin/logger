package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T){
	logger :=NewFileLogger(LogLevelDebug,"F:/logs/","test")
	logger.Debug("user id[%d] is come from china",323)
	logger.Debug("user id[%d] is come from china",323)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()

}


func TestConsoleLogger(t *testing.T){
	logger :=NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id[%d] is come from china",323)
	logger.Debug("user id[%d] is come from china",323)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()

}