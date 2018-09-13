package types

import (
	"errors"
	"reflect"
)

// Call - call to specified method
type Call struct {
	method     func()
	methodHash string
}

// NewCall - initialize new instance of Call struct
func NewCall(method func()) (*Call, error) {
	if reflect.ValueOf(method).IsNil() { // Check for nil method
		return &Call{}, errors.New("nil call") // Return error
	}

	call := Call{method, ""} // Init call

	return &call, nil // Return initialized call instance
}

// Run - attempt to run specified call
func (*Call) Run() error {
	return nil
}
