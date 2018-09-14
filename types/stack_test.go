package types

import (
	"errors"
	"reflect"
	"testing"
)

/* BEGIN EXPORTED METHODS */

// TestNewStack - test functionality of stack initializer
func TestNewStack(t *testing.T) {
	stack := NewStack() // Initialize stack

	if reflect.ValueOf(stack).IsNil() {
		t.Errorf(errors.New("nil stack").Error()) // Log found error
		t.FailNow()                               // Panic
	}

	t.Logf("initialized stack %p successfully", stack) // Log success
}

// TestStackAddCall - test functionality of addcall stack method
func TestStackAddCall(t *testing.T) {
	stack := NewStack() // Initialize stack

	if reflect.ValueOf(stack).IsNil() {
		t.Errorf(errors.New("nil stack").Error()) // Log found error
		t.FailNow()                               // Panic
	}

	t.Logf("initialized stack %p successfully", stack) // Log success

	call, err := NewCall(callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = stack.AddCall(call) // Attempt to add call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("added call %v successfully", call)
}

func TestRunStack(t *testing.T) {
	stack := NewStack() // Initialize stack

	if reflect.ValueOf(stack).IsNil() {
		t.Errorf(errors.New("nil stack").Error()) // Log found error
		t.FailNow()                               // Panic
	}

	call, err := NewCall(callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = stack.AddCall(call) // Attempt to add call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = stack.Run() // Attempt to run stack

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("ran stack %v successfully", stack)
}

/* END EXPORTED METHODS */
