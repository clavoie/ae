package ae

import "google.golang.org/appengine/datastore"

// KeyStatus represents the disposition of a user supplied entitiy
// key string
type KeyStatus byte

const (
	// KeyStatusNew indicates that the use supplied key is empty,
	// and the user is attempting to save a new object to the db
	KeyStatusNew KeyStatus = iota
	// KeyStatusExisting indicates that the user supplied key is
	// valid, and represents an existing object
	KeyStatusExisting
	// KeyStatusErr indicates the the user supplied key is invalid,
	// and could not be decoded
	KeyStatusErr
)

// IKeyStatus is a dependency which knows how to convert user
// supplied key strings into datastore.Keys and KeyStatus codes
type IKeyStatus interface {
	Decode(string) (KeyStatus, *datastore.Key)
}

// keyStatusImpl is an implementation of IKeyStatus
type keyStatusImpl struct{}

// newKeyStatus returns a new instance of IKeyStatus
func newKeyStatus() IKeyStatus {
	return new(keyStatusImpl)
}

func (ks *keyStatusImpl) Decode(userKey string) (KeyStatus, *datastore.Key) {
	if userKey == "" {
		return KeyStatusNew, nil
	}

	key, err := datastore.DecodeKey(userKey)
	if err == nil {
		return KeyStatusExisting, key
	}

	return KeyStatusErr, nil
}
