package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// rope2 is a struct that has a chain with 10 links. chain[0] is the head, chain[9] is the tail.
type rope2 struct {
	l         zerolog.Logger
	chain     [10][2]int
	tailTrace map[int]struct{}
}

func newRopeTwo(l zerolog.Logger) *rope2 {
	localLogger := l.With().Str("part", "rope two").Logger()
	r := &rope2{
		l: localLogger,
		chain: [10][2]int{
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
		},
		tailTrace: make(map[int]struct{}),
	}

	r.tailTrace[coordToBinary(r.chain[9][0], r.chain[9][1])] = struct{}{}

	return r
}

func (r *rope2) moveHead(dir int, dist int) error {
	r.l.Debug().Msgf("MOV HEAD %s %d", dirToString(dir), dist)
	for i := 0; i < dist; i++ {
		switch dir {
		case up:
			r.chain[0] = [2]int{r.chain[0][0] - 1, r.chain[0][1]}
		case right:
			r.chain[0] = [2]int{r.chain[0][0], r.chain[0][1] + 1}
		case down:
			r.chain[0] = [2]int{r.chain[0][0] + 1, r.chain[0][1]}
		case left:
			r.chain[0] = [2]int{r.chain[0][0], r.chain[0][1] - 1}
		default:
			return fmt.Errorf("this should not have happened, got direction %#v", dir)
		}

		err := r.moveChain()
		if err != nil {
			return errors.Wrapf(err, "moving chain in dir %s by %d", dirToString(dir), dist)
		}
	}

	r.l.Debug().Msgf("\n\n%s\n", r.visual())

	return nil
}

func (r *rope2) moveChain() error {
	for i := 0; i < len(r.chain)-1; i++ {
		moved, err := r.moveNext(i, i+1)
		if err != nil {
			return errors.Wrapf(err, "moving chain from %d to %d had an issue", i, i+1)
		}

		// this chain did not move, therefore none of the others following it will either.
		if moved == [2]int{0, 0} {
			// early exit
			return nil
		}
	}
	// we moved all the chain pieces without issue
	return nil
}

func (r *rope2) moveNext(curr, next int) ([2]int, error) {
	upDownDistance := r.chain[curr][0] - r.chain[next][0]
	leftRightDistance := r.chain[curr][1] - r.chain[next][1]
	if upDownDistance < -2 || upDownDistance > 2 ||
		leftRightDistance < -2 || leftRightDistance > 2 {
		return [2]int{}, fmt.Errorf("up-down or left-right distance is too big! udd: %d, lrd: %d, head: %v, tail: %v",
			upDownDistance, leftRightDistance, r.chain[curr], r.chain[next])
	}

	moveBy := [2]int{0, 0}
	switch [2]int{upDownDistance, leftRightDistance} {
	case [2]int{-2, 0}:
		// tail is directly above
		moveBy = [2]int{-1, 0}
	case [2]int{-2, -1}, [2]int{-2, -2}, [2]int{-1, -2}:
		// tail is in the upper right quadrant
		moveBy = [2]int{-1, -1}
	case [2]int{0, -2}:
		// tail is directly to the right
		moveBy = [2]int{0, -1}
	case [2]int{1, -2}, [2]int{2, -2}, [2]int{2, -1}:
		// tail is in the lower right quadrant
		moveBy = [2]int{1, -1}
	case [2]int{2, 0}:
		// tail is directly below
		moveBy = [2]int{1, 0}
	case [2]int{2, 1}, [2]int{2, 2}, [2]int{1, 2}:
		// tail is in the bottom left quadrant
		moveBy = [2]int{1, 1}
	case [2]int{0, 2}:
		// tail is directly to the left
		moveBy = [2]int{0, 1}
	case [2]int{-1, 2}, [2]int{-2, 2}, [2]int{-2, 1}:
		// tail is in the upper left quadrant
		moveBy = [2]int{-1, 1}
	default:
		return [2]int{}, nil
	}

	r.chain[next] = [2]int{r.chain[next][0] + moveBy[0], r.chain[next][1] + moveBy[1]}
	if next == 9 {
		r.l.Debug().Msgf("LNK Tail moved to %v", r.chain[next])
		r.tailTrace[coordToBinary(r.chain[next][0], r.chain[next][1])] = struct{}{}
	}

	return moveBy, nil
}

func (r *rope2) placesBeen() int {
	return len(r.tailTrace)
}

func (r *rope2) visual() string {
	var sb strings.Builder
	minx, miny, maxx, maxy := 0, 0, 0, 0
	for _, link := range r.chain {
		if link[0] < minx {
			minx = link[0]
		}

		if link[0] > maxx {
			maxx = link[0]
		}

		if link[1] < miny {
			miny = link[1]
		}
		if link[1] > maxy {
			maxy = link[1]
		}
	}

	grid := make(map[int]map[int]string)

	for row := minx; row <= maxx; row++ {
		grid[row] = make(map[int]string)
		for col := miny; col <= maxy; col++ {
			grid[row][col] = "."
		}
	}

	// mark start
	grid[0][0] = "s"

	// mark head
	grid[r.chain[0][0]][r.chain[0][1]] = "H"

	for i := 9; i > 0; i-- {
		grid[r.chain[i][0]][r.chain[i][1]] = strconv.Itoa(i)
	}

	for row := minx; row <= maxx; row++ {
		for col := miny; col <= maxy; col++ {
			sb.WriteString(grid[row][col])
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
