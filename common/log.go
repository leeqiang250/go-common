package common

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"io/ioutil"
	"log"
)

var (
	Trace   *log.Logger
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
)

func logfileInit(
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	fatalHandle io.Writer) {

	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	Debug = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	Fatal = log.New(fatalHandle, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInit(logfile string, level int) (ok bool) {
	var traceHandle, debugHandle io.Writer

	// I just hard code it,bite me
	infoLogHandler := &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    500,
		MaxBackups: 20,
		MaxAge:     7,
		Compress:   false,
	}

	wfLogHandler := &lumberjack.Logger{
		Filename:   logfile + ".wf",
		MaxSize:    500,
		MaxBackups: 20,
		MaxAge:     7,
		Compress:   false,
	}

	if level >= 16 {
		traceHandle = infoLogHandler
		debugHandle = infoLogHandler
	} else if level >= 8 {
		traceHandle = ioutil.Discard
		debugHandle = infoLogHandler
	} else {
		traceHandle = ioutil.Discard
		debugHandle = ioutil.Discard
	}

	logfileInit(traceHandle,
		debugHandle,
		infoLogHandler,
		wfLogHandler,
		wfLogHandler,
		wfLogHandler)

	return true
}