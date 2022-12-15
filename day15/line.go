package day15

import (
	"sort"

	"github.com/pkg/errors"
)

const (
	lineAngled orientation = iota
	lineHorizontal
	lineVertical
)

type orientation int

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

func (l line) isCoordInLine(c coordinate) bool {
	if l.orientation == lineHorizontal {
		if l.rowCol == c[1] {
			return l.start[0] <= c[0] && l.end[0] >= c[0]
		}

		return false
	}

	if l.rowCol == c[0] {
		return l.start[1] <= c[1] && l.end[1] >= c[1]
	}

	return false
}

type lines []line

func (l lines) Len() int {
	return len(l)
}

func (l lines) Less(i, j int) bool {
	if l[i].orientation != l[j].orientation {
		return l[i].orientation < l[j].orientation
	}

	if l[i].orientation == lineVertical {
		return l[i].start[1] < l[j].start[1]
	}

	return l[i].start[0] < l[j].start[0]
}

func (l lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
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

func reduceLines(ls lines) (lines, error) {
	for {
		sort.Sort(ls)

		startLength := ls.Len()
		i := 0

		for {
			if len(ls)-1 <= i {
				break
			}
			a := ls[i]
			b := ls[i+1]

			m, err := mergeLines(a, b)
			if err != nil {
				// do not change the slice, move on to the next one
				i++
				continue
			}
			// m is now a and b merged
			ls[i] = m
			_, ls, err = pluck(ls, i+1)
		}

		if ls.Len() == startLength {
			break
		}
	}

	return ls, nil
}
