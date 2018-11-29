package logger


import (
	"runtime"
	"path"
	"time"
	"fmt"
)

type LogData struct{
	Message string
	TimeStr string
	LevelStr string
	Filename string
	FuncName string
	LineNo int
	WarnAndFatal bool
}

func GetLineInfo()(fileName string,funcName string,lineNo int){
	pc ,file,line,ok :=runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}


func writeLog(level int,format string,args ...interface{}) *LogData{
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")

	levelStr := getLevelText(level)
	fileName ,funcName ,lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format,args...)

	logData :=&LogData{
		Message : msg,
		TimeStr:nowStr,
		LevelStr:levelStr,
		Filename:fileName,
		FuncName:funcName,
		LineNo:lineNo,
		WarnAndFatal:false,
	}
	if level==LogLevelError||level==LogLevelWarn||level==LogLevelFatal{
		logData.WarnAndFatal = true
	}
	// fmt.Fprintf(file,"%s %s (%s:%s:%d) %s\n",nowStr,levelStr,fileName,funcName,lineNo,msg)
	return logData
}