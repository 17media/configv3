package configv3

import (
	"regexp"
)

// regexChan ties up a Regexp with a channel, chan string
type regexChan struct {
	regex *regexp.Regexp
	ch    *chan ModifiedFile
}

// Client provides a interface for configuration service
type Client interface {
	// ConfigInfo returns the ConfigInfo protobuf struct that
	// contains infomation w.r.t to repo and last commit
	ConfigInfo() ConfigInfo

	// Get request the content of file at path. If the path doesn't
	// exists then ok returns false
	Get(path string) ([]byte, error)

	// AddListener returns a chan that when any file that matches
	// the pathRegEx changes, the ModifiedFile will be sent over
	// through the channel
	AddListener(pathRegEx *regexp.Regexp) *chan ModifiedFile

	// Watch watches change of a file and invokes callback. It also invokes callback for the
	// first time and return error if there's one.
	Watch(path string, callback func([]byte) error, errChan chan<- error) error

	// RemoveListener remove a listener channel
	RemoveListener(ch *chan ModifiedFile)

	// List lists content under certain path
	List(path string) (map[string][]byte, error)

	// Stop is to stop client
	Stop() error

	// BumpSum bump stat
	BumpSum(key string, val float64)

	// GetRoot get root path
	GetRoot() string
}
