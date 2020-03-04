package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type ApplicationLogger interface {
	Trace(t interface{})
	Info(i interface{})
	Warning(w interface{})
	Error(e interface{})
	Fatal(f interface{})
}

type appLogger struct {
	TraceLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
}

var singletonLogger ApplicationLogger

var once sync.Once

var (
	ErrInvalidLogLevel = errors.New("Invalid Log level.  Must be a value from 0 to 4.")
)

//NewLogger Retrieves a thread-safe singleton logger that logs at the given level:
// 0 - Trace or higher
// 1 - Info or higher (probably what you want)
// 2 - Warning or higher
// 3 - Error or higher
// 4 - Fatal or higher
func NewLogger(logLevel int) (ApplicationLogger, error) {
	var err error = nil
	once.Do(func() {
		if logLevel < 0 || logLevel > 4 {
			err = ErrInvalidLogLevel
			return
		}

		logger := &appLogger{
			log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile),
			log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile),
			log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile),
			log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile),
			log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile),
		}

		if logLevel < 1 {
			logger.SetTraceLogger(os.Stdout)
		}

		if logLevel < 2 {
			logger.SetInfoLogger(os.Stdout)
		}

		if logLevel < 3 {
			logger.SetWarningLogger(os.Stdout)
		}

		if logLevel < 4 {
			logger.SetErrorLogger(os.Stdout)
		}

		if logLevel < 5 {
			logger.SetFatalLogger(os.Stdout)
		}

		singletonLogger = logger
		return
	})

	if err != nil {
		return nil, err
	}

	return singletonLogger, nil
}

func (al *appLogger) SetTraceLogger(w io.Writer) {
	if w != nil {
		al.TraceLogger = log.New(w, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		return
	}

	al.TraceLogger = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func (al *appLogger) SetInfoLogger(w io.Writer) {
	if w != nil {
		al.InfoLogger = log.New(w, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		return
	}

	al.InfoLogger = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func (al *appLogger) SetWarningLogger(w io.Writer) {
	if w != nil {
		al.WarningLogger = log.New(w, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		return
	}

	al.WarningLogger = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func (al *appLogger) SetErrorLogger(w io.Writer) {
	if w != nil {
		al.ErrorLogger = log.New(w, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		return
	}

	al.ErrorLogger = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func (al *appLogger) SetFatalLogger(w io.Writer) {
	if w != nil {
		al.FatalLogger = log.New(w, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		return
	}

	al.FatalLogger = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func (al *appLogger) Trace(t interface{}) {
	al.TraceLogger.Printf("%v", t)
}

func (al *appLogger) Info(i interface{}) {
	al.InfoLogger.Printf("%v", i)
}

func (al *appLogger) Warning(w interface{}) {
	al.WarningLogger.Printf("%v", w)
}

func (al *appLogger) Error(e interface{}) {
	al.ErrorLogger.Printf("%v", e)
}

func (al *appLogger) Fatal(f interface{}) {
	al.FatalLogger.Printf("%v", f)
}
