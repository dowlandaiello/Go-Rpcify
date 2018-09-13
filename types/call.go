package types

import (
	"errors"
	"reflect"

	"github.com/mitsukomegumi/Go-Rpcify/common"
)

// Call - call to specified method
type Call struct {
	Method     func() `json:"method"` // Used to call method
	MethodHash string `json:"hash"`   // Used to identify call

	Endpoint string `json:"endpoint"` // Used for calls to rpc
}

/* BEGIN EXPORTED METHODS */

// NewCall - initialize new instance of Call struct
func NewCall(method func(), endpoint string) (*Call, error) {
	if reflect.ValueOf(method).IsNil() { // Check for nil method
		return &Call{}, errors.New("nil call") // Return error
	}

	methodHash := ""

	call := Call{method, methodHash, endpoint} // Init call

	byteValue, err := common.ToBytes(call) // Attempt to encode

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	call.MethodHash = common.SHA3String(byteValue) // Calculate method hash

	if call.Endpoint == "" { // Check for nil endpoint
		call.Endpoint = common.RootCallEndpoint + "/" + call.MethodHash // Set endpoint
	}

	return &call, nil // Return initialized call instance
}

// Run - attempt to run specified call
func (call *Call) Run() error {
	call.Method() // Run method

	return nil // No error occurred, return nil
}

/* END EXPORTED METHODS */
