package stack

import "container/list"

// Stack represents a classic stack. The zero value for Stack is an empty stack ready to use.
type Stack struct {
	list *list.List
}

// New returns an initialized Stack.
func New() *Stack {
	return &Stack{list: list.New()}
}

// Len returns the length of stack
func (s *Stack) Len() int {
	return s.list.Len()
}

// Peek returns the top element of Stack without removing it; or nil if the Stack is empty.
func (s *Stack) Peek() interface{} {
	if s.Len() == 0 {
		return nil
	}

	return s.list.Front().Value
}

// Pop removes the top element of Stack and return it; or returns nil if Stack is empty.
func (s *Stack) Pop() interface{} {
	e := s.list.Front()
	if e == nil {
		return nil
	}

	s.list.Remove(e)
	return e.Value
}

// Push pushes to the top of Stack.
func (s *Stack) Push(value interface{}) interface{} {
	e := s.list.PushFront(value)
	return e.Value
}
