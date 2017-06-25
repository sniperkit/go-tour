package stack

import "errors"

// Represents a Stack.
type Stack []interface{}

// Creates a new Stack
func New() Stack {
	return make(Stack, 0)
}

// Pushes all of v onto the stack.
func (s *Stack) Push(v ...interface{}) {
	for _, i := range v {
		*s = append(*s, i)
	}
}

// Pops an element from the stack. An attempt to pop
// an empty stack is an error.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return 0, errors.New("cannot pop an empty stack")
	}

	l := len(*s)

	ret := (*s)[l-1]
	*s = (*s)[:l-1]

	return ret, nil
}

// Is true if the stack is empty.
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
