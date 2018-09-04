package ae

import "google.golang.org/appengine/datastore"

// Key is a wrapper around the app engine's datastore *Key struct
type Key interface {
	IntId() int64

	ToDatastore() *datastore.Key
}

// keyImpl is an implementation of IKey
type keyImpl struct {
	key *datastore.Key
}

// newKey returns a new IKey initialized from an existing key
func newKey(key *datastore.Key) Key {
	return &keyImpl{
		key: key,
	}
}

func (k *keyImpl) IntId() int64                { return k.key.IntID() }
func (k *keyImpl) ToDatastore() *datastore.Key { return k.key }
