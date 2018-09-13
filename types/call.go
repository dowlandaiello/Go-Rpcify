package types

import (
	"errors"
	"reflect"

	"github.com/mitsukomegumi/Go-Rpcify/common"
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

	byteValue, err := common.ToBytes(method) // Attempt to encode

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	methodHash := common.SHA3String(byteValue)

	call := Call{method, methodHash} // Init call

	return &call, nil // Return initialized call instance
}

// Run - attempt to run specified call
func (*Call) Run() error {
	return nil // No error occurred, return nil
}
