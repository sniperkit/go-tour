package stack

import "errors"

type Stack []interface{}

func New() Stack {
	return make(Stack, 0)
}

func (s *Stack) Push(v ...interface{}) {
	for _, i := range v {
		*s = append(*s, i)
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return 0, errors.New("cannot pop an empty stack")
	}

	l := len(*s)

	ret := (*s)[l-1]
	*s = (*s)[:l-1]

	return ret, nil
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
