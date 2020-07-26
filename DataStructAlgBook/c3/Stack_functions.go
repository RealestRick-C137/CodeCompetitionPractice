package main

import (
  "fmt"
)

// Returns top element of stack without removing
func (s *Stack) Top() int {
  return (*s)[len(*s) - 1 ]
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(element int) {
	*s = append(*s, element) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Pop all element of the stack
func (s *Stack) Dump() {
  for s.IsEmpty() == false {
    val, _ := s.Pop()
    fmt.Println(val)
  }
}
