package common

import (
	"errors"
	"testing"
)

/*
	BEGIN CRYPTO UNIT TESTING METHODS
*/

// TestSHA3String - test functionality of SHA3String method
func TestSHA3String(t *testing.T) {
	test := SHA3String([]byte("test")) // Hash test string

	if test != "NvAoWAuwLMgnKpoCD0IA40bidq5mTkXugHRVdOL1q4A=" { // Check for matching hash
		t.Errorf(errors.New("invalid hash").Error()) // Log found error
		t.FailNow()                                  // Panic
	}

	t.Logf("found hash %s", test) // Log success
}

// TestSHA3Bytes - test functionality of SHA3Bytes method
func TestSHA3Bytes(t *testing.T) {
	test := SHA3Bytes([]byte("test")) // Hash test string

	if test == nil { // Check for nil hash
		t.Errorf(errors.New("invalid hash").Error()) // Log found error
		t.FailNow()                                  // Panic
	}

	t.Logf("found hash %b", test) // Log success
}

/*
	END CRYPTO UNIT TESTING METHODS
*/

/*
	BEGIN CONVERSION UNIT TESTING METHODS
*/

// TestToBytes - test functionality of ToBytes method
func TestToBytes(t *testing.T) {
	test, err := ToBytes("test") // Test ToBytes

	if err != nil { // Check for error
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found byte value %b", test) // Log success
}

/*
	END CONVERSION UNIT TESTING METHODS
*/
