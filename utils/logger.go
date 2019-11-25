package utils

import (
	"io"
	"log"
)

var (
	// LogTrace is for debug only
	LogTrace *log.Logger
	
	LogInfo *log.Logger
	
	LogWarn *log.Logger
	
	LogError *log.Logger
)

func LogInit(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warnHandle io.Writer,
	errorHandle io.Writer) {

		LogTrace = log.New(traceHandle, "TRACE:", log.Ldate|log.Ltime|log.Lshortfile)
		LogInfo = log.New(infoHandle, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
		LogWarn = log.New(warnHandle, "WARN:", log.Ldate|log.Ltime|log.Lshortfile)
		LogError = log.New(errorHandle, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}

