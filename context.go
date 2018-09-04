package ae

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

// Context is an injectable wrapper around golang.org/x/net/context.
type Context interface {
	// Context returns the original Context object.
	Context() context.Context
}

// contextImpl is an implementation of the Context interface.
type contextImpl struct {
	context context.Context
}

// NewContext returns a new instance of Context.
func NewContext(r *http.Request) Context {
	return &contextImpl{appengine.NewContext(r)}
}

// NewTransactionContext returns a new Context instance that
// wraps a datastore.RunInTransaction context
func NewTransactionContext(ctx context.Context) Context {
	return &contextImpl{ctx}
}

func (ci *contextImpl) Context() context.Context {
	return ci.context
}
