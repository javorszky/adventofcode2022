package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Sponsored by Dwayne Johnson

const (
	matEmpty material = iota
	matRock
	matSand
	matAir
	matEntry

	maskRow int = 0b1111111111100000000000
	maskCol int = 0b0000000000011111111111
)

type material int

type rock struct {
	grid                           map[int]material
	minRow, minCol, maxRow, maxCol int
}

func newRock() rock {
	return rock{
		grid:   make(map[int]material),
		minRow: 0,
		minCol: 0,
		maxRow: 0,
		maxCol: 0,
	}
}

func xyToBinary(row, col int) int {
	return row<<11 | col
}

func binaryToXY(in int) (int, int) {
	row := in & maskRow
	col := in & maskCol
	return row >> 11, col
}

// addRocks takes the input, parses it, and creates the map that has the air
// and rock coordinates.
//
// From AOC D14: " ...reports the x,y coordinates that form the shape of the
// path, where x represents distance to the right and y represents distance
// down."
//
// Based on the above:
// - x is col, increase to the right
// - y is row, increase down
func (r *rock) addRocks(lines []string) error {
	minCol, minRow, maxCol, maxRow := 1<<32, 0, 0, 0
	rocks := make([]int, 0)
	for _, in := range lines {
		fmt.Printf("adding line\n%s\n\n", in)
		vertices := strings.Split(in, " -> ")
		coordList := make([][2]int, len(vertices))
		for i := 0; i < len(vertices); i++ {
			pts := strings.Split(vertices[i], ",")
			if len(pts) != 2 {
				return fmt.Errorf("parsing '%s' into pair of coords, is not 2 parts", vertices[i])
			}

			col, err := strconv.Atoi(pts[0])
			if err != nil {
				return errors.Wrapf(err, "converting %s col into int", pts[0])
			}

			row, err := strconv.Atoi(pts[1])
			if err != nil {
				return errors.Wrapf(err, "converting %s row into int", pts[1])
			}

			// expand map
			if row < minRow {
				minRow = row
			}

			if col < minCol {
				minCol = col
			}

			if row > maxRow {
				maxRow = row
			}

			if col > maxCol {
				maxCol = col
			}

			coordList[i] = [2]int{row, col}
		}

		for i := 0; i < len(coordList)-1; i++ {
			list, err := generateConnectors(coordList[i], coordList[i+1])
			if err != nil {
				return errors.Wrapf(err, "generating connectors for %v -> %v", coordList[0], coordList[1])
			}
			fmt.Printf(" -- connectors: %v\n", list)

			for _, cc := range list {
				fmt.Printf(" --x- adding rock to %d, %d\n", cc[0], cc[1])
				rocks = append(rocks, xyToBinary(cc[0], cc[1]))
			}
		}
	}

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			r.grid[xyToBinary(row, col)] = matAir
		}
	}

	for _, pebble := range rocks {
		r.grid[pebble] = matRock
	}

	r.grid[xyToBinary(0, 500)] = matEntry

	r.minRow = minRow
	r.maxRow = maxRow
	r.minCol = minCol
	r.maxCol = maxCol

	return nil
}

func (r *rock) String() string {
	var sb strings.Builder
	draw := " "

	for row := r.minRow; row <= r.maxRow; row++ {
		for col := r.minCol; col <= r.maxCol; col++ {
			switch r.grid[xyToBinary(row, col)] {
			case matAir:
				draw = "."
			case matRock:
				draw = "#"
			case matSand:
				draw = "~"
			case matEntry:
				draw = "+"
			default:
				draw = "?"

			}
			sb.WriteString(draw)
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func generateConnectors(a, b [2]int) ([][2]int, error) {
	c := make([][2]int, 0)

	// are we doing x or y?
	if a[0] == b[0] {
		// we're doing a y diff
		small, large := a[1], b[1]
		if large < small {
			small, large = large, small
		}

		for i := small; i <= large; i++ {
			c = append(c, [2]int{a[0], i})
		}

		return c, nil
	}

	if a[1] != b[1] {
		return nil, errors.New("it would be diagonal, what u doin bae")
	}

	small, large := a[0], b[0]
	if large < small {
		small, large = large, small
	}

	for i := small; i <= large; i++ {
		c = append(c, [2]int{i, a[1]})
	}
	return c, nil
}
