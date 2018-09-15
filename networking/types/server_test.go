package types

import (
	"errors"
	"reflect"
	"testing"
)

// TestNewServer - test functionality of server initializer
func TestNewServer(t *testing.T) {
	server := NewServer("test") // Initialize server

	if reflect.ValueOf(server).IsNil() {
		t.Errorf(errors.New("nil server").Error()) // Log found error
		t.FailNow()                                // Panic
	}

	t.Logf("initialized server %v", server) // Log success
}

// TestStartServer - test functionality of StartServer method
func TestStartServer(t *testing.T) {
	server := NewServer("test") // Initialize server

	if reflect.ValueOf(server).IsNil() {
		t.Errorf(errors.New("nil server").Error()) // Log found error
		t.FailNow()                                // Panic
	}

	go server.StartServer() // Start server

	t.Logf("started server %v", server) // Log success
}
