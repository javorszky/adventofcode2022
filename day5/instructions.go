package day5

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// instructions is a slice of [3]int. That [3]int contains 3 pieces of data: how many to move at [0], source stack
// designation [1], and destination stack designation [2].
type instructions [][3]int

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
