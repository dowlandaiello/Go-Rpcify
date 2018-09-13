package types

import (
	"errors"
	"reflect"
)

// Stack - struct used for referencing/running multiple calls at once
type Stack struct {
	Calls []*Call `json:"calls"` // Used for referencing and running Stack calls
}

/* BEGIN EXPORTED METHODS */

// NewStack - initialize new instance of Stack type
func NewStack() *Stack {
	return &Stack{} // Return initialized Stack
}

// AddCall - attempt to append call to stack
func (stack *Stack) AddCall(call *Call) error {
	if reflect.ValueOf(call).IsNil() { // Check for nil call
		return errors.New("invalid call") // Return found error
	}

	(*stack).Calls = append((*stack).Calls, call) // Append call

	return nil // No error occurred, return nil
}

// Run - iterate through stack, run each call
func (stack *Stack) Run() error {
	if reflect.ValueOf(stack).IsNil() { // Check for errors
		return errors.New("nil stack") // Return found error
	}

	for x := 0; x != len(stack.Calls); x++ { // Iterate through stack
		stack.Calls[x].Run() // Run call
	}

	return nil // No error occurred, return nil
}

/* END EXPORTED METHODS */
