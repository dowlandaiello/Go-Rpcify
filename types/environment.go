package types

// Environment - stack-based container capable of holding memory-only runtime values
type Environment struct {
	stack *Stack // Fetch reference to
}

// NewEnvironment - initialize and return new instance of Environment struct
func NewEnvironment() *Environment { // TODO: finished initializer
	envStack := NewStack() // Initialize stack

	env := Environment{stack: envStack} // Init new environment

	return &env // Return initialized instance
}
