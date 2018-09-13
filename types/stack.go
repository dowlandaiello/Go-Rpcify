package types

// Stack - struct used for referencing/running multiple calls at once
type Stack struct {
	Calls []*Call `json:"calls"` // Used for referencing and running Stack calls
}

/* BEGIN EXPORTED METHODS */

// NewStack - initialize new instance of Stack type
func NewStack() *Stack {
	return &Stack{} // Return initialized Stack
}

/* END EXPORTED METHODS */
