package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 5).Int("part", 1).Logger()

	gog, err := inputs.GroupByBlankLines("day5/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	st := makeStacks(gog[0])

	insts, err := makeInstructions(gog[1])
	if err != nil {
		localLogger.Err(err).Msgf("parsing instructions")
		os.Exit(1)
	}

	for _, inst := range insts {
		for i := 0; i < inst[0]; i++ {
			st[inst[2]].Push(st[inst[1]].Pop())
		}
	}

	var sb strings.Builder
	for i := 1; i < 10; i++ {
		sb.WriteString(st[i].Pop())
	}
	// code goes here

	solution := sb.String()
	s := localLogger.With().Str("solution", solution).Logger()
	s.Info().Msgf("After shuffling the boxes, the top ones are '%s'", solution)
}

// instructions is a slice of [3]int. That [3]int contains 3 pieces of data: how many to move at [0], source stack
// designation [1], and destination stack designation [2].
type instructions [][3]int

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

func makeInstructions(in []string) (instructions, error) {
	replacer := strings.NewReplacer("move ", "", " from ", " ", " to ", " ")
	instruct := make(instructions, len(in))
	for i, line := range in {
		numbers := replacer.Replace(line)
		each := strings.Split(numbers, " ")
		m, err := strconv.Atoi(each[0])
		if err != nil {
			return nil, errors.Wrapf(err, "converting %s to int on line %d: %s", each[0], i, line)
		}

		f, err := strconv.Atoi(each[1])
		if err != nil {
			return nil, errors.Wrapf(err, "converting %s to int on line %d: %s", each[1], i, line)
		}

		t, err := strconv.Atoi(each[2])
		if err != nil {
			return nil, errors.Wrapf(err, "converting %s to int on line %d: %s", each[2], i, line)
		}

		instruct[i] = [3]int{m, f, t}
	}
	return instruct, nil
}
