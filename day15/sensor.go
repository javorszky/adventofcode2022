package day15

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

const (
	lineAngled orientation = iota
	lineHorizontal
	lineVertical
)

type orientation int

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
	}

	if l.orientation == lineVertical {
		return line{}, fmt.Errorf("line is vertical for row %d: bounds %v, %v", row, a, b)
	}

	return l, nil
}

type grid struct {
	sensors []sensor
}

func (g *grid) addSensor(s sensor) {
	g.sensors = append(g.sensors, s)
}

func (g *grid) sensorsExcludingRow(row int) []sensor {
	listOfSensors := make([]sensor, 0)

	for _, s := range g.sensors {
		if s.rowInExclusion(row) {
			listOfSensors = append(listOfSensors, s)
		}
	}

	return listOfSensors
}

type line struct {
	start, end  coordinate
	orientation orientation
	rowCol      int
}

func (l line) Len() int {
	if l.orientation == lineVertical {
		return l.end[1] - l.start[1] + 1
	}
	return l.end[0] - l.start[0] + 1
}

// newLine creates a new line defined by two coordinates. It has to be either vertical or horizontal,
// which means either the first or second coordinates NEED to match.
func newLine(a, b coordinate) (line, error) {
	if a[0] != b[0] && a[1] != b[1] {
		return line{}, errors.New("the two points do not make for a horizontal or vertical line")
	}

	start := a
	end := b

	if a[0] == b[0] {
		// vertical
		if start[1] > end[1] {
			start, end = end, start
		}

		return line{
			start:       start,
			end:         end,
			orientation: lineVertical,
			rowCol:      a[0],
		}, nil
	}

	// horizontal
	if start[0] > end[0] {
		start, end = end, start
	}

	return line{
		start:       start,
		end:         end,
		orientation: lineHorizontal,
		rowCol:      a[1],
	}, nil
}

func mergeLines(a, b line) (line, error) {
	if a.orientation != b.orientation {
		return line{}, errors.New("lines are different orientation")
	}

	if a.rowCol != b.rowCol {
		return line{}, errors.New("lines are on different planes")
	}

	if a.orientation == lineHorizontal {
		// horizontal, meaning [1] will stay the same
		if a.start[0] > b.start[0] {
			a, b = b, a
		}

		if b.start[0]-a.end[0] > 1 {
			// too far apart, can't merge
			return line{}, errors.New("don't touch, can't merge")
		}

		end := b.end
		if a.end[0] > end[0] {
			end = a.end
		}

		return line{
			start:       a.start,
			end:         end,
			orientation: a.orientation,
			rowCol:      a.rowCol,
		}, nil
	}

	// vertical, meaning [0] will stay the same
	if a.start[1] > b.start[1] {
		a, b = b, a
	}

	if b.start[1]-a.end[1] > 1 {
		// too far apart, can't merge
		return line{}, errors.New("don't touch, can't merge")
	}

	end := b.end
	if a.end[1] > end[1] {
		end = a.end
	}

	return line{
		start:       a.start,
		end:         end,
		orientation: a.orientation,
		rowCol:      a.rowCol,
	}, nil
}

func reduceLines(lines []line) ([]line, error) {
	sort.Slice(lines, func(i, j int) bool {
		if lines[i].orientation != lines[j].orientation {
			return lines[i].orientation < lines[j].orientation
		}

		if lines[i].orientation == lineVertical {
			return lines[i].start[1] < lines[j].start[1]
		}

		return lines[i].start[0] < lines[j].start[0]
	})

	i := 0
	for {
		if len(lines)-1 <= i {
			break
		}
		a := lines[i]
		b := lines[i+1]

		m, err := mergeLines(a, b)
		if err != nil {
			// do not change the slice, move on to the next one
			i++
			continue
		}
		// m is now a and b merged
		lines[i] = m
		_, lines, err = pluck(lines, i+1)
	}

	return lines, nil
}

func pluck[T any](sl []T, idx int) (T, []T, error) {
	var thing T
	if idx >= len(sl) {
		return thing, nil, errors.New("index out of bounds")
	}

	thing = sl[idx]
	ls := append(sl[:idx], sl[idx+1:]...)
	return thing, ls, nil
}

func newSensor(own, closest coordinate) sensor {
	return sensor{
		self:              own,
		closestBeacon:     closest,
		manhattanDistance: manhattanDistance(own, closest),
	}
}

func newGrid() grid {
	return grid{sensors: make([]sensor, 0)}
}

func absDiff(a, b int) int {
	c := a - b
	if c < 0 {
		return b - a
	}
	return c
}

func manhattanDistance(a, b coordinate) int {
	return absDiff(a[0], b[0]) + absDiff(a[1], b[1])
}
