package day5

import (
	"fmt"
)

// stacks is a map of *stack with addresses 1-9, so we can use the move commands easily.
type stacks map[int]*stack

// stack represents a stack of crates. The zeroth element is the one on the bottom, the last element is the one on top.
// This is a FILO queue, so push / pop always operates at the end of the slice.
type stack struct {
	elements []string
}

func (s *stack) Push(in string) string {
	s.elements = append(s.elements, in)
	return in
}

func (s *stack) Pop() string {
	out := s.elements[len(s.elements)-1:]
	s.elements = s.elements[:len(s.elements)-1]
	return out[0]
}

func (s *stack) PopMany(n int) ([]string, error) {
	if n > len(s.elements) {
		return nil, fmt.Errorf("trying to remove too many elements. Want %d, have %d", n, len(s.elements))
	}
	out := s.elements[len(s.elements)-n:]
	s.elements = s.elements[:len(s.elements)-n]
	return out, nil
}

func (s *stack) PushMany(in []string) {
	s.elements = append(s.elements, in...)
}

func (s *stack) String() string {
	return fmt.Sprintf("%v", s.elements)
}

func NewStack() *stack {
	return &stack{elements: make([]string, 0)}
}

// makeStacks takes the first half of day 5 input 1 and turns it into a collection of stacks.
//
//	[S] [J] [C]     [F] [C]     [D] [G]
//	 1   5   9   13  17  21  25  29  33
func makeStacks(lines []string) stacks {
	// set up the container
	st := make(stacks)
	for i := 1; i <= 9; i++ {
		st[i] = NewStack()
	}

	// create the list of columns where boxes are
	cols := []int{1, 5, 9, 13, 17, 21, 25, 29, 33}

	// remove the last line with the indexes
	lines = lines[:len(lines)-1]

	// add the boxes to the stacks line by line from the bottom up, if there are boxes.
	for i := len(lines) - 1; i >= 0; i-- {
		for j, c := range cols {
			box := string(lines[i][c])
			if box != " " {
				st[j+1].Push(box)
			}
		}
	}

	return st
}
