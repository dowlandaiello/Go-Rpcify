package types

import "testing"

/* BEGIN INTERNAL METHODS */

// callTestMethod
func callTestMethod() {
	print("test")
}

/* END INTERNAL METHODS */

/* BEGIN EXPORTED METHODS */

// TestNewCall - test functionality of Call initializer method
func TestNewCall(t *testing.T) {
	call, err := NewCall(callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("initialized call %s successfully", call.MethodHash) // Log success
}

/* END EXPORTED METHODS */
