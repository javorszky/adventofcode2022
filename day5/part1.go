package day5

import (
	"fmt"
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 5).Int("part", 1).Logger()

	gog, err := inputs.GroupByBlankLines("day5/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	ssdf := "A-Z"

	fmt.Printf("%o", []byte(ssdf))

	// code goes here
	fmt.Printf("top half goes here\n\n%#v", gog[0])

	solution := 2
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("-- change this for the part 1 message -- %d", solution)
}

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

func (s *stack) String() string {
	return fmt.Sprintf("%v", s.elements)
}

func NewStack() *stack {
	return &stack{elements: make([]string, 0)}
}
