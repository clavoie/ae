package ae

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// ILogger represents a general interface around the logging
// implementation
type ILogger interface {
	Debugf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Infof(fmt string, args ...interface{})
	Warningf(fmt string, args ...interface{})
}

// appEngineLogger represents a logging implementation which logs
// to the Google App Engine
type appEngineLogger struct {
	context context.Context
}

// newLogger returns a new app engine logger and maps it to
// ILogger
//
// This needs to be public because it is accessed by the main program
func NewLogger(r *http.Request) ILogger {
	return &appEngineLogger{
		context: appengine.NewContext(r),
	}
}

func (ael *appEngineLogger) Debugf(fmt string, args ...interface{}) {
	log.Debugf(ael.context, fmt, args...)
}

func (ael *appEngineLogger) Errorf(fmt string, args ...interface{}) {
	log.Errorf(ael.context, fmt, args...)
}

func (ael *appEngineLogger) Infof(fmt string, args ...interface{}) {
	log.Infof(ael.context, fmt, args...)
}

func (ael *appEngineLogger) Warningf(fmt string, args ...interface{}) {
	log.Warningf(ael.context, fmt, args...)
}
