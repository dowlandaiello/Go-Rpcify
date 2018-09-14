package types

import (
	"errors"
	"reflect"
	"testing"
)

/* BEGIN EXPORTED METHODS */

// TestNewEnvironment - test functionality of environment initializer
func TestNewEnvironment(t *testing.T) {
	env := NewEnvironment("test") // Initialize new environment

	if reflect.ValueOf(env).IsNil() {
		t.Errorf(errors.New("nil environment").Error()) // Log found error
		t.FailNow()                                     // Panic
	}

	t.Logf("initialized new environment %v", env) // Log success
}

// TestAddStack - test functionality of AddStack method
func TestAddStack(t *testing.T) {
	env := NewEnvironment("test") // Initialize new environment

	if reflect.ValueOf(env).IsNil() { // Check for nil env
		t.Errorf(errors.New("nil environment").Error()) // Log found error
		t.FailNow()                                     // Panic
	}

	call, err := NewCall(callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	stack := NewStack("test") // Initialize stack

	err = stack.AddCall(call) // Add call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = env.AddStack(stack) // Attempt to add stack

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("added stack %v successfully", stack) // Log success
}

// TestAddCall - test functionality of AddCall method
func TestAddCall(t *testing.T) {
	env := NewEnvironment("test") // Initialize new environment

	if reflect.ValueOf(env).IsNil() { // Check for nil env
		t.Errorf(errors.New("nil environment").Error()) // Log found error
		t.FailNow()                                     // Panic
	}

	t.Logf("initialized new environment %v", env) // Log success

	call, err := NewCall(callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("initialized call %s successfully", call.MethodHash) // Log success

	err = env.AddCall(call) // Attempt to add call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
	}

	t.Logf("added call %v successfully", call) // Log success
}

/* END EXPORTED METHODS */
