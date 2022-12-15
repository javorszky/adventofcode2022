package day15

import (
	"fmt"

	"github.com/pkg/errors"
)

// coordinate holds a pair of ints, in col / row format.
// x, y =>
//
//	x is column, horizontal, right is plus
//	y is row, vertical, down is plus
type coordinate [2]int

type sensor struct {
	self              coordinate
	closestBeacon     coordinate
	manhattanDistance int
}

func (s sensor) rowInExclusion(row int) bool {
	return s.self[1]-s.manhattanDistance <= row && row <= s.self[1]+s.manhattanDistance
}

func (s sensor) lineForRow(row int) (line, error) {
	if !s.rowInExclusion(row) {
		return line{}, fmt.Errorf("row %d is not in exclusion zone for this sensor", row)
	}

	d := s.manhattanDistance - absDiff(row, s.self[1])
	a, b := coordinate{s.self[0] - d, row}, coordinate{s.self[0] + d, row}

	l, err := newLine(a, b)
	if err != nil {
		return line{}, errors.Wrapf(err, "creating line for %v, %v", a, b)
	}

	if a == b && l.orientation == lineVertical {
		l.orientation = lineHorizontal
		l.rowCol = row
	}

	if l.orientation == lineVertical {
		return line{}, fmt.Errorf("line is vertical for row %d: bounds %v, %v", row, a, b)
	}

	return l, nil
}

func newSensor(own, closest coordinate) sensor {
	return sensor{
		self:              own,
		closestBeacon:     closest,
		manhattanDistance: manhattanDistance(own, closest),
	}
}
