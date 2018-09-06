package ae

import "google.golang.org/appengine/datastore"

// Db is a wrapper around app engine's datastore package.
type Db interface {
	// Add inserts a value into the datastore, creating a new incomplete
	// key, inserting the value, and returning the new complete key.
	Add(string, interface{}) (Key, error)

	// Get returns an object from the database. The first return value
	// indicates if the object exists or not
	Get(Key, interface{}) (bool, error)

	// NewKey returns a new key from a kind, intID, and parent key.
	// An incomplete key can be returned by passing 0 and nil for the
	// final two arguments. The kind is required
	NewKey(string, int64, Key) Key

	// NewQuery returns a new Query from a db kind string
	NewQuery(string) Query

	// Put does an upsert on an entity
	Put(Key, interface{}) (Key, error)
}

// dbImpl is an implementation of Db
type dbImpl struct {
	context Context
}

// NewDb creates a new instance of IDb
func NewDb(context Context) Db {
	return &dbImpl{
		context: context,
	}
}

func (db *dbImpl) Add(kind string, src interface{}) (Key, error) {
	var err error
	key := datastore.NewIncompleteKey(db.context.Context(), kind, nil)
	key, err = datastore.Put(db.context.Context(), key, src)

	if err != nil {
		return nil, err
	}

	return NewKey(key), nil
}

func (db *dbImpl) Get(key Key, dst interface{}) (bool, error) {
	err := datastore.Get(db.context.Context(), key.ToDatastore(), dst)

	if err == datastore.ErrNoSuchEntity {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *dbImpl) NewKey(kind string, intID int64, parent Key) Key {
	var parentKey *datastore.Key

	if parent != nil {
		parentKey = parent.ToDatastore()
	}

	aeKey := datastore.NewKey(db.context.Context(), kind, "", intID, parentKey)
	return NewKey(aeKey)
}

func (db *dbImpl) NewQuery(kind string) Query {
	return newQuery(db.context, kind)
}

func (db *dbImpl) Put(key Key, src interface{}) (Key, error) {
	aeKey, err := datastore.Put(db.context.Context(), key.ToDatastore(), src)

	if err != nil {
		return nil, err
	}

	return NewKey(aeKey), nil
}
