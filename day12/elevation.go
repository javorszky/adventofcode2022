package day12

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

type elevationMap struct {
	grid  map[int]map[int]int32
	start [2]int
	goal  [2]int
}

func parseInput(gog []string) *elevationMap {
	start, goal := [2]int{}, [2]int{}
	m := make(map[int]map[int]int32)
	for row, line := range gog {
		m[row] = make(map[int]int32)
		for col, char := range line {
			m[row][col] = char
			if char == 83 {
				start = [2]int{row, col}
			}
			if char == 69 {
				goal = [2]int{row, col}
			}
		}
	}
	em := &elevationMap{
		grid:  m,
		start: start,
		goal:  goal,
	}

	return em
}
