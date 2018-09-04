package ae

import "google.golang.org/appengine/datastore"

// Query is a wrapper around app engine's datastore Query object
type Query interface {
	// Filter applies a filter to a query, returning the new
	// query
	Filter(string, interface{}) Query

	// GetAll runs the query and returns all results and keys for
	// matching records. For details please see:
	//
	// https://cloud.google.com/appengine/docs/standard/go/datastore/reference#Query
	GetAll(interface{}) ([]Key, error)

	// KeysOnly returns a new query that only returns keys and
	// ignores the destination value in GetAll
	KeysOnly() Query

	// Limit limits the number of results that are returned
	Limit(int) Query
}

// queryImpl is an implementation of IQuery
type queryImpl struct {
	context Context
	query   *datastore.Query
}

// newQuery creates a new IQuery from a database kind string
func newQuery(context Context, kind string) Query {
	return newQueryInternal(context, datastore.NewQuery(kind))
}

// newQueryInternal creates a new IQuery from an existing datastore.Query.
// This func should only be called by this file
func newQueryInternal(context Context, query *datastore.Query) Query {
	return &queryImpl{
		context: context,
		query:   query,
	}
}

func (q *queryImpl) Filter(filterStr string, value interface{}) Query {
	newQuery := q.query.Filter(filterStr, value)
	return newQueryInternal(q.context, newQuery)
}

func (q *queryImpl) GetAll(dst interface{}) ([]Key, error) {
	keys, err := q.query.GetAll(q.context.Context(), dst)

	if err != nil {
		return nil, err
	}

	ikeys := make([]Key, len(keys))
	for index, key := range keys {
		ikeys[index] = newKey(key)
	}

	return ikeys, nil
}

func (q *queryImpl) KeysOnly() Query {
	newQuery := q.query.KeysOnly()
	return newQueryInternal(q.context, newQuery)
}

func (q *queryImpl) Limit(limit int) Query {
	newQuery := q.query.Limit(limit)
	return newQueryInternal(q.context, newQuery)
}
