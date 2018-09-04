package ae

import "google.golang.org/appengine"

// AppEngine is a wrapper around the google.golang.org/appengine
// package.
type AppEngine interface {
	// AppId returns the application ID for the current application.
	// See: https://cloud.google.com/appengine/docs/standard/go/reference#AppID
	AppId() string

	// IsDevAppServer reports if the current application is running in the
	// development App Server.
	// See: https://cloud.google.com/appengine/docs/standard/go/reference#IsDevAppServer
	IsDevAppServer() bool

	// ModuleName returns the name of the current module.
	// See: https://cloud.google.com/appengine/docs/standard/go/reference#ModuleName
	ModuleName() string
}

// appEngineImpl is an implementation of the AppEngine interface.
type appEngineImpl struct {
	context Context
}

// NewAppEngine returns a new instance of an AppEngine implementation.
func NewAppEngine(context Context) AppEngine {
	return &appEngineImpl{context}
}

func (aei *appEngineImpl) AppId() string {
	return appengine.AppID(aei.context.Context())
}

func (aei *appEngineImpl) IsDevAppServer() bool {
	return appengine.IsDevAppServer()
}

func (aei *appEngineImpl) ModuleName() string {
	return appengine.ModuleName(aei.context.Context())
}
