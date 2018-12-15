package types

import "testing"

/* BEGIN INTERNAL METHODS */

// callTestMethod - method for testing calls
func callTestMethod() (string, error) {
	print("-- CALL -- test\n") // Log success

	return "-- CALL -- test", nil // No error occurred, return nil
}

/* END INTERNAL METHODS */

/* BEGIN EXPORTED METHODS */

// TestNewCall - test functionality of Call initializer method
func TestNewCall(t *testing.T) {
	call, err := NewCall(callTestMethod, callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("initialized call %v successfully with endpoint %s", call, call.Endpoint) // Log success
}

// TestRunCall - test functionality of call run method
func TestRunCall(t *testing.T) {
	call, err := NewCall(callTestMethod, callTestMethod, "") // Init call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("initialized call %s successfully", call.MethodHash) // Log success

	output, err := call.Run() // Attempt to run call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
	}

	t.Logf("ran call %s successfully. Encountered output: %s", call.MethodHash, output) // Log success
}

/* END EXPORTED METHODS */
