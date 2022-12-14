package day8

import (
	"sort"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func newForest(coords []string, l zerolog.Logger) (forest, error) {
	localLogger := l.With().Str("part", "forest").Logger()
	sizeY := len(coords)
	sizeX := len(coords[0])
	trees := make(map[int]map[int]int)
	for i, row := range coords {
		// row is y coordinate
		trees[i] = make(map[int]int)
		for j, col := range row {
			h, err := strconv.Atoi(string(col))
			if err != nil {
				return forest{}, errors.Wrapf(err, "conversion to int from char %d/%s", col, string(col))
			}
			trees[i][j] = h
		}
	}

	return forest{
		l:     localLogger,
		sizeX: sizeX,
		sizeY: sizeY,
		trees: trees,
	}, nil
}

type forest struct {
	l            zerolog.Logger
	sizeX, sizeY int
	trees        map[int]map[int]int
}

func (f forest) adjacentCoords(x, y int) [][2]int {
	// grab all 4 coordinates
	coords := [][2]int{
		{x, y + 1},
		{x, y - 1},
		{x + 1, y},
		{x - 1, y},
	}
	out := make([][2]int, 0)
	for _, c := range coords {
		// filter out coordinates that are out of bounds:
		// x or y is below 0, or
		// x or y is above the max size
		if c[0] < 0 || c[0] >= f.sizeX || c[1] < 0 || c[1] >= f.sizeY {
			continue
		}

		out = append(out, c)
	}

	return out
}

// countVisible will count the trees from the edges in a straight line
//
// 0 ----- x ----->
// |
// |
// y
// |
// |
// v
func (f forest) countVisible() int {
	visible := make(map[uint16]struct{})
	max := 0
	height := 0

	// from left edge towards right
	for row := 0; row < f.sizeY; row++ {
		f.l.Debug().Msgf("Row %d ⮕", row)
		for col := 0; col < f.sizeX; col++ {
			height = f.trees[row][col]
			bcoord := coordToBinary(row, col)

			// left edge gets col 0 visible
			if col == 0 {
				max = height
				visible[bcoord] = struct{}{}
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				continue
			}

			if height > max {
				max = height
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				visible[bcoord] = struct{}{}
				continue
			}
			f.l.Debug().Msgf("%d - %d%s", max, height, "")
		}
	}

	// from right edge towards left
	for row := 0; row < f.sizeY; row++ {
		f.l.Debug().Msgf("Row %d ⬅", row)
		for col := f.sizeX - 1; col >= 0; col-- {
			height = f.trees[row][col]
			bcoord := coordToBinary(row, col)

			// right edge gets last col visible
			if col == f.sizeX-1 {
				max = height
				visible[bcoord] = struct{}{}
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				continue
			}

			if height > max {
				max = height
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				visible[bcoord] = struct{}{}
				continue
			}
			f.l.Debug().Msgf("%d - %d%s", max, height, "")
		}
	}

	// from bottom edge up
	f.l.Debug().Msgf("===> Bottom edge towards up")
	for col := 0; col < f.sizeX; col++ {
		f.l.Debug().Msgf("Column %d ⬆", col)
		for row := f.sizeY - 1; row >= 0; row-- {
			height = f.trees[row][col]
			bcoord := coordToBinary(row, col)

			// bottom gets last row visible always
			if row == f.sizeY-1 {
				max = height
				visible[bcoord] = struct{}{}
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				continue
			}

			if height > max {
				max = height
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				visible[bcoord] = struct{}{}
				continue
			}
			f.l.Debug().Msgf("%d - %d%s", max, height, "")
		}
	}

	// from top edge down
	f.l.Debug().Msgf("===> Top edge towards down")
	for col := 0; col < f.sizeX; col++ {
		f.l.Debug().Msgf("Column %d ⬇", col)
		for row := 0; row < f.sizeY; row++ {
			height = f.trees[row][col]
			bcoord := coordToBinary(row, col)

			// top gets row 0 visible always
			if row == 0 {
				max = height
				visible[bcoord] = struct{}{}
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				continue
			}

			if height > max {
				max = height
				f.l.Debug().Msgf("%d - %d%s", max, height, "*")
				visible[bcoord] = struct{}{}
				continue
			}
			f.l.Debug().Msgf("%d - %d%s", max, height, "")
		}
	}

	return len(visible)
}

func (f forest) mostScenic() int {
	scores := make([]int, (f.sizeX-2)*(f.sizeY-2))
	i := 0
	// HELL YEAH CPU GO BRRRR
	for row := 1; row < f.sizeY-1; row++ {
		for col := 1; col < f.sizeX-1; col++ {
			scores[i] = f.scenicScore(row, col)
			i++
		}
	}

	sort.Ints(scores)

	return scores[len(scores)-1]
}

func (f forest) scenicScore(row, col int) int {
	h := f.trees[row][col]
	dirs := [4][2]int{
		{-1, 0}, // row -1, col same => up
		{0, -1}, // row same, col -1 => left
		{1, 0},  // row +1, col same => down
		{0, 1},  // row same, col +1 => right
	}
	visible := [4]int{}

	f.l.Debug().Msg("")
	f.l.Debug().Msgf("Calc for [%d, %d]: %d", row, col, h)

	for i, d := range dirs {
		f.l.Debug().Msgf("Direction %#v", d)
		row_, col_ := row, col
		dv := 0
		for {
			// advance the check coords by current direction
			row_, col_ = row_+d[0], col_+d[1]
			h_, ok := f.trees[row_][col_]
			if !ok {
				f.l.Debug().Msg("edge")
				break
			}
			dv++
			f.l.Debug().Msgf("-- %d - %d", h_, h)
			if h_ >= h {
				f.l.Debug().Msgf("big")
				break
			}
		}
		f.l.Debug().Msgf("seen %d for dir %v", dv, d)
		visible[i] = dv
	}
	product := 1
	for _, n := range visible {
		product = product * n
	}

	f.l.Debug().Msgf("product is %d for visibles %v", product, visible)

	return product
}

func coordToBinary(x, y int) uint16 {
	return uint16(x<<8 | y)
}

func foo() {
	_ = map[int]map[int]int{
		0: {0: 3, 1: 0, 2: 3, 3: 7, 4: 3},
		1: {0: 2, 1: 5, 2: 5, 3: 1, 4: 2},
		2: {0: 6, 1: 5, 2: 3, 3: 3, 4: 2},
		3: {0: 3, 1: 3, 2: 5, 3: 4, 4: 9},
		4: {0: 3, 1: 5, 2: 3, 3: 9, 4: 0}}
}
