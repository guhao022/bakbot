package axiom

import (
	"fmt"
)

const (
	DefaultAdapter = `shell`
)

// Adapter interface
type Adapter interface {
	Name() string
	// New() (Adapter, error)
	Run() error
	Stop() error

	Receive(*Message) error
	Send(*Response, ...string) error
	Emote(*Response, ...string) error
	Reply(*Response, ...string) error
	Topic(*Response, ...string) error
	Play(*Response, ...string) error

	String() string
}

type adapter struct {
	name     string
	newFunc  func(*Robot) (Adapter, error)
	sendChan chan *Response
	recvChan chan *Message
}

// AvailableAdapters is a map of registered adapters
var AvailableAdapters = map[string]adapter{}

// NewAdapter creates a new initialized adapter
func NewAdapter(robot *Robot) (Adapter, error) {
	name := DefaultAdapter
	if _, ok := AvailableAdapters[name]; !ok {
		return nil, fmt.Errorf("%s is not a registered adapter", name)
	}

	adapter, err := AvailableAdapters[name].newFunc(robot)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}

// RegisterAdapter registers an adapter
func RegisterAdapter(name string, newFunc func(*Robot) (Adapter, error)) {
	AvailableAdapters[name] = adapter{
		name:    name,
		newFunc: newFunc,
	}
}

// BasicAdapter declares common functions shared by all adapters
type BasicAdapter struct {
	*Robot
}

// SetRobot sets the adapter's Robot
func (a *BasicAdapter) SetRobot(r *Robot) {
	a.Robot = r
}

func (a *BasicAdapter) String() string {
	return DefaultAdapter
}

