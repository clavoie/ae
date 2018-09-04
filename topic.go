package ae

// ITopic is a wrapper around a google cloud PubSub topic
type ITopic interface {
	Exists() (bool, error)

	Publish()

	Stop()
}
