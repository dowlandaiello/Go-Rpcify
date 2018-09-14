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

/* END EXPORTED METHODS */
