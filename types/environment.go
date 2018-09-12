package types

// Environment - stack-based container capable of holding memory-only runtime values
type Environment struct {
}

// NewEnvironment - initialize and return new instance of Environment struct
func NewEnvironment() *Environment { // TODO: finished initializer
	env := Environment{}

	return &env
}
