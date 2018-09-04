package ae

import "github.com/clavoie/di"

// NewDiDefs returns the dependency injection definitions for this package
func NewDiDefs() []*di.Def {
	return []*di.Def{
		{NewAppEngine, di.PerHttpRequest},
		{NewContext, di.PerHttpRequest},
		{NewDb, di.PerHttpRequest},
		{NewEnv, di.Singleton},
		{NewKeyDecoder, di.Singleton},
		{NewTaskQueue, di.PerHttpRequest},
	}
}
