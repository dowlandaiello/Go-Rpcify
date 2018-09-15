package upnp

import "testing"

// TestDiscoverGateway - test functionality of gateway discovery
func TestDiscoverGateway(t *testing.T) {
	_, err := GetGateway() // Attempt to fetch gateway device

	if err != nil {
		t.Errorf(err.Error()) // If error occurs, print error to testing console
		t.FailNow()           // Panic
	}
}

// TestForwardPort - test functionality of port forwarding
func TestForwardPort(t *testing.T) {
	err := ForwardPort(8080) // Attempt to forward port 3000

	if err != nil {
		t.Errorf(err.Error()) // If error occurs, print error to testing console
		t.FailNow()           // Panic
	}
}

func TestRemovePortForward(t *testing.T) {
	err := RemovePortForward(8080) // Attempt to remove forward on port 3000

	if err != nil {
		t.Errorf(err.Error()) // If error occurs, print error to testing console
		t.FailNow()           // Panic
	}
}
