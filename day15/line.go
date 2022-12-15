package day15

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
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

func cutLines(origin line, cutSet lines, l zerolog.Logger) (lines, error) {
	originalParts := lines{origin}
	filteredCutset := make(lines, 0)

	for _, cs := range cutSet {
		if cs.orientation == origin.orientation && cs.rowCol == origin.rowCol {
			filteredCutset = append(filteredCutset, cs)
		}
	}

	// if len(filteredCutset) == 0 {
	// 	l.Debug().Msgf("filtered cutset is zero")
	// 	return originalParts, nil
	// }

	l.Debug().Msgf("filtered cutset is %v", filteredCutset)

	idx := 0
	if origin.orientation == lineVertical {
		l.Debug().Msgf("lines are vertical (y (idx 1) moves, x (idx 0) stationary)")
		idx = 1
	}

	allNewParts := make(lines, 0)
	newParts := originalParts

	for _, fc := range filteredCutset {

		for _, op := range newParts {
			ops, ope, fcs, fce := op.start[idx], op.end[idx], fc.start[idx], fc.end[idx]

			l.Debug().Msgf("Op %d -> %d, FC %d -> %d", ops, ope, fcs, fce)

			switch {
			case ope < fcs || ops > fce:
				l.Debug().Msgf("do not overlap, op is added back")
				// do not overlap, continue
				// Fs--Fe ... OPs--OPe
				// OPs--OPe ... Fs--Fe
				// ... do nothing
				// add op back to the new parts unchanged
				newParts = append(newParts, op)

			case fcs <= ops && fce >= ope:
				// cutset envelops, op disappears
				// Fs--OPs--OPe--Fe
				// ... remove op from slice
				// do not add op to the new parts
				l.Debug().Msgf("cutset envelops, op disappears")

			case ops < fcs && fce < ope:
				l.Debug().Msgf("cutset contained, op broken into two")
				// cutset contained, op broken into two
				// OPs--Fs--Fe--OPe
				// OPs--Fs-1 ... Fe+1--OPe
				if origin.orientation == lineHorizontal {
					ls, err := newLine(coordinate{ops, origin.rowCol}, coordinate{fcs - 1, origin.rowCol})
					if err != nil {
						return nil, errors.Wrapf(err, "cutset contained, left new line, op broken into two. OP %v, FC %v", op, fc)
					}

					le, err := newLine(coordinate{fce + 1, origin.rowCol}, coordinate{ope, origin.rowCol})
					if err != nil {
						return nil, errors.Wrapf(err, "cutset contained, right new line, op broken into two. OP %v, FC %v", op, fc)
					}

					newParts = append(newParts, ls, le)
					continue
				}

				ls, err := newLine(coordinate{origin.rowCol, ops}, coordinate{origin.rowCol, fcs - 1})
				if err != nil {
					return nil, errors.Wrapf(err, "cutset contained, top new line, op broken into two. OP %v, FC %v", op, fc)
				}

				le, err := newLine(coordinate{origin.rowCol, fce + 1}, coordinate{origin.rowCol, ope})
				if err != nil {
					return nil, errors.Wrapf(err, "cutset contained, bottom new line, op broken into two. OP %v, FC %v", op, fc)
				}

				newParts = append(newParts, ls, le)

			case fcs <= ops && fce < ope:
				l.Debug().Msgf("cutset overlaps on left")
				// cutset overlaps on left
				// Fs--OPs-Fe---OPe
				// Fe+1--OPe
				if origin.orientation == lineHorizontal {
					lr, err := newLine(coordinate{fce + 1, origin.rowCol}, coordinate{ope, origin.rowCol})
					if err != nil {
						return nil, errors.Wrapf(err, "cutset overlaps on left. OP %v, FC %v", op, fc)
					}
					newParts = append(newParts, lr)
					continue
				}

				lr, err := newLine(coordinate{origin.rowCol, fce + 1}, coordinate{origin.rowCol, ope})
				if err != nil {
					return nil, errors.Wrapf(err, "cutset overlaps on top. OP %v, FC %v", op, fc)
				}
				newParts = append(newParts, lr)
			case ops < fcs && ope <= fce:
				l.Debug().Msgf("cutset overlaps on right")
				// cutset overlaps on right
				// OPs--Fs-OPe--Fe
				// OPs--Fs-1
				if origin.orientation == lineHorizontal {
					ll, err := newLine(coordinate{ops, origin.rowCol}, coordinate{fcs - 1, origin.rowCol})
					if err != nil {
						return nil, errors.Wrapf(err, "cutset overlaps on right, OP %v, FC %v", op, fc)
					}
					newParts = append(newParts, ll)
					continue
				}

				ll, err := newLine(coordinate{origin.rowCol, ops}, coordinate{origin.rowCol, fcs - 1})
				if err != nil {
					return nil, errors.Wrapf(err, "cutset overlaps on top, OP %v, FC %v", op, fc)
				}
				newParts = append(newParts, ll)
			default:
				return nil, fmt.Errorf("none of the above matched for op %v and fc %v", op, fc)
			}
		}

	}

	allNewParts = append(allNewParts, newParts...)

	m, err := reduceLines(allNewParts)
	if err != nil {
		return nil, errors.Wrapf(err, "tried to merge allnewparts: %v", allNewParts)
	}

	return m, nil
}
