package logs

import (
	"io/ioutil"
	"log"
	"os"
)

type LogState int

const (
	DebugMode = iota
	WarningMode
	InfoMode
	Quiet
)

var (
	Mode    LogState
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Initialize() {
	switch Mode {
	case DebugMode:
		log.SetOutput(os.Stdout)
		Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case WarningMode:
		log.SetOutput(os.Stdout)
		Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case InfoMode:
		log.SetOutput(os.Stdout)
		Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(ioutil.Discard, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(ioutil.Discard, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	case Quiet:
		log.SetOutput(ioutil.Discard)
		Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(ioutil.Discard, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(ioutil.Discard, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		break
	}
}
