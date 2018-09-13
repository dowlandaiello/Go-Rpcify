package types

// Call - call to specified method
type Call struct {
	method func()
}

// NewCall - initialize new instance of Call struct
func NewCall(method func()) *Call {
	call := Call{method}

	return &call
}
