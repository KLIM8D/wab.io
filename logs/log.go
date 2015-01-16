package logs

import (
	"io"
	"log"
)

var (
	Output  io.Writer
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Initialize() {
	log.SetOutput(Output)
	Trace = log.New(Output, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(Output, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(Output, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(Output, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
