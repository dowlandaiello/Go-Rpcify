package types

import "testing"

/* BEGIN INTERNAL METHODS */

// callTestMethod - method for testing calls

func callTestMethod() {
	print("-- CALL -- test") // Log success
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

// TestRunCall - test functionality of call run method
func TestRunCall(t *testing.T) {
	call, err := NewCall(callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("initialized call %s successfully", call.MethodHash) // Log success

	err = call.Run() // Attempt to run call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
	}

	t.Logf("ran call %s successfully", call.MethodHash)
}

/* END EXPORTED METHODS */
