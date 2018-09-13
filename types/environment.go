package types

// Environment - stack-based container capable of holding memory-only runtime values
type Environment struct {
	Stacks []*Stack `json:"stacks"` // Fetch reference to Environment stacks
}

// NewEnvironment - initialize and return new instance of Environment struct
func NewEnvironment() *Environment { // TODO: finished initializer
	envStack := NewStack() // Initialize default stack

	env := Environment{[]*Stack{envStack}} // Init new environment

	return &env // Return initialized instance
}
