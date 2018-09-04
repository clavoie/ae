package ae

import "google.golang.org/appengine/datastore"

// Key is a wrapper around the app engine's datastore *Key struct
type Key interface {
	// IntId returns the Key's integer ID.
	// See: https://cloud.google.com/appengine/docs/standard/go/datastore/reference#Key
	// IntID()
	IntId() int64

	// ToDatastore returns the app engine's datastore representation of the Key
	ToDatastore() *datastore.Key
}

// keyImpl is an implementation of Key
type keyImpl struct {
	key *datastore.Key
}

// NewKey returns a new Key initialized from an existing key
func NewKey(key *datastore.Key) Key {
	return &keyImpl{
		key: key,
	}
}

func (k *keyImpl) IntId() int64                { return k.key.IntID() }
func (k *keyImpl) ToDatastore() *datastore.Key { return k.key }
