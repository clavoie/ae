package ae

import "os"

// nameEnv is the environment variable name for which the value of the
// running environment is stored
const nameEnv = "RSPP_ENV"

// Env is a wrapper arnound environment variables provided to the
// program
type Env interface {
	IsProd() bool
	IsStaging() bool
}

// envImpl is an implementation of Env
type envImpl struct {
	isProd    bool
	isStaging bool
}

// NewEnv returns a new instance of Env
func NewEnv() Env {
	return &envImpl{
		isProd:    os.Getenv(nameEnv) == "prod",
		isStaging: os.Getenv(nameEnv) == "staging",
	}
}

func (e *envImpl) IsProd() bool {
	return e.isProd
}

func (e *envImpl) IsStaging() bool {
	return e.isStaging
}