package day12

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rs/zerolog"
)

// char to int32 conversion table.
// a, 97
// b, 98
// c, 99
// d, 100
// e, 101
// f, 102
// g, 103
// h, 104
// i, 105
// j, 106
// k, 107
// l, 108
// m, 109
// n, 110
// o, 111
// p, 112
// q, 113
// r, 114
// s, 115
// t, 116
// u, 117
// v, 118
// w, 119
// x, 120
// y, 121
// z, 122
//
// S, 83
// E, 69

const (
	colOffset = 0b0000000000000001
	rowOffset = 0b0000000100000000
	colMask   = 0b0000000011111111
	rowMask   = 0b1111111100000000
)

type elevationMap struct {
	l      zerolog.Logger
	grid   map[int]int32
	width  int
	height int
	start  int
	goal   int
}

// coordToBinary will turn row and col into a single number.
//
// row is max 41, so up to 64, 2^6
// col is max 144, so up to 256, so 2^8
func coordToBinary(row, col int) int {
	return row<<8 | col
}

func binaryToCoord(coord int) (int, int) {
	col := coord & colMask
	row := (coord & rowMask) >> 8
	return row, col
}

func parseInput(gog []string, l zerolog.Logger) *elevationMap {
	start, goal := 0, 0
	m := make(map[int]int32)
	for row, line := range gog {
		for col, char := range line {
			cb := coordToBinary(row, col)
			m[cb] = char
			// start
			if char == 83 {
				m[cb] = 97
				start = cb
			}
			// end
			if char == 69 {
				m[cb] = 122
				goal = cb
			}
		}
	}
	em := &elevationMap{
		grid:   m,
		width:  len(gog[0]),
		height: len(gog),
		start:  start,
		goal:   goal,
		l:      l.With().Str("module", "elevation map").Logger(),
	}

	return em
}

func (em *elevationMap) String() string {
	var sb strings.Builder
	for row := 0; row < em.height; row++ {
		for col := 0; col < em.width; col++ {
			v, ok := em.grid[coordToBinary(row, col)]
			if !ok {
				sb.WriteString(" ")
				continue
			}
			sb.WriteString(string(v))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (em *elevationMap) shortestRoute() [][]int {
	visitedInSteps := make(map[int]int)
	var goThere func(int, int, int32, []int, int) [][]int

	goThere = func(coord, previousCoord int, previousElevation int32, route []int, depth int) [][]int {
		row, col := binaryToCoord(coord)
		pRow, pCol := binaryToCoord(previousCoord)
		currentElevation := em.grid[coord]

		em.l.Debug().Msgf("checking [%d, %d](%d) -> [%d, %d](%d) {%d}", pRow, pCol, previousElevation, row, col, currentElevation, depth)

		route = append(route, coord)
		if coord == em.goal {
			em.l.Debug().Msgf("!!! - we have found a route! %v", route)
			// we have found a way!
			return [][]int{route}
		}

		routes := make([][]int, 0)

		// we have not been here before, or we have but this is now a shorter route
		visitedInSteps[coord] = depth

		for _, c := range []int{
			moveRight(coord),
			moveDown(coord),
			moveLeft(coord),
			moveUp(coord),
		} {
			if c == previousCoord {
				em.l.Debug().Msgf("<<<< no backtrack")
				continue
			}

			if err := em.canGoThere(c, depth+1, currentElevation, visitedInSteps); err != nil {
				drow, dcol := binaryToCoord(c)
				em.l.Debug().Msgf("-- can't go to %d, %d: %s", drow, dcol, err.Error())
				continue
			}

			routeCopy := make([]int, len(route))
			for k, v := range route {
				routeCopy[k] = v
			}

			routes = append(routes, goThere(c, coord, currentElevation, routeCopy, depth+1)...)
		}

		return routes
	}

	em.l.Debug().Msgf("okay, starting at the goal point...")
	// start at goal which is at elevation z, which is code point 122
	routes := goThere(em.start, em.start, em.grid[em.start], []int{}, 0)

	em.l.Debug().Msgf("routes we got back are\n%#v", routes)
	sort.Slice(routes, func(i, j int) bool {
		return len(routes[i]) < len(routes[j])
	})

	return routes
}

func (em *elevationMap) canGoThere(coord, depth int, elevation int32, visited map[int]int) error {
	v, ok := em.grid[coord]
	row, col := binaryToCoord(coord)
	if !ok {
		// tile does not exist
		return fmt.Errorf("tile does not exist at %d, %d", row, col)
	}

	diff := v - elevation
	if diff < -1 || diff > 1 {
		// tile too tall / deep to go there
		return fmt.Errorf("elevation is too big at %d, %d, (%d %s -> %d %s)", row, col, elevation, string(elevation), v, string(v))
	}

	seen, ok := visited[coord]
	if ok && seen <= depth {
		// we've been there, and already found a shorter route
		return fmt.Errorf("been there, found a shorter route: %d, %d (%d vs depth %d)", row, col, seen, depth)
	}

	return nil
}

func moveUp(c int) int {
	return c - rowOffset
}

func moveDown(c int) int {
	return c + rowOffset
}

func moveLeft(c int) int {
	return c - colOffset
}

func moveRight(c int) int {
	return c + colOffset
}
