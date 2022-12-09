package day9

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/pkg/errors"
)

const (
	up = iota
	right
	down
	left
)

// rope keeps track of the current position of the rope's head, tail, and the history of where the tail's been. The
// first element of the position array is for the up / down, where positive is up, the second one is left / right where
// positive is right.
//
// There's also a bounds, which keeps track of the maximum values for each direction. Indexes and directions:
// 0: up
// 1: right
// 2: down
// 3: left
type rope struct {
	l               zerolog.Logger
	head            [2]int
	tail            [2]int
	tailTrace       map[int]struct{}
	tailTraceString map[string]struct{}
}

func newRope(l zerolog.Logger) *rope {
	localLogger := l.With().Str("part", "rope").Logger()
	r := &rope{
		l:               localLogger,
		head:            [2]int{0, 0},
		tail:            [2]int{0, 0},
		tailTrace:       make(map[int]struct{}),
		tailTraceString: make(map[string]struct{}),
	}

	r.tailTrace[coordToBinary(r.tail[0], r.tail[1])] = struct{}{}
	r.tailTraceString[coordToString(r.tail[0], r.tail[1])] = struct{}{}

	return r
}

func (r *rope) moveHead(dir int, dist int) error {
	r.l.Debug().Msgf("%s %d", dirToString(dir), dist)
	for i := 0; i < dist; i++ {
		switch dir {
		case up:
			r.head = [2]int{r.head[0] + 1, r.head[1]}
		case right:
			r.head = [2]int{r.head[0], r.head[1] + 1}
		case down:
			r.head = [2]int{r.head[0] - 1, r.head[1]}
		case left:
			r.head = [2]int{r.head[0], r.head[1] - 1}
		default:
			return fmt.Errorf("this should not have happened, got direction %#v", dir)
		}

		err := r.moveTail()
		if err != nil {
			return errors.Wrapf(err, "moving head in dir %s by %d", dirToString(dir), dist)
		}
		r.l.Debug().Msgf("%v  /  %v", r.head, r.tail)
	}

	return nil
}

// moveTail will check whether the tail needs to be moved, and if it does, moves it to the appropriate location.
//
// Scenarios
// 1. H T - xdistance =  0, ydistance = -2 // tail is more positive to the right
// 2. T H - xdistance =  0, ydistance =  2 // head is more positive to the right
// 3. H
//
// # T   - xdistance =  2, ydistance =  0 // head is more positive to up
//
// 4. T
//
// # H   - xdistance = -2, ydistance =  0 // tail is more positive to up
func (r *rope) moveTail() error {
	upDownDistance := r.head[0] - r.tail[0]
	leftRightDistance := r.head[1] - r.tail[1]
	if upDownDistance < -2 || upDownDistance > 2 ||
		leftRightDistance < -2 || leftRightDistance > 2 {
		return fmt.Errorf("up-down or left-right distance is too big! udd: %d, lrd: %d, head: %v, tail: %v",
			upDownDistance, leftRightDistance, r.head, r.tail)
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
		return nil
	}

	r.tail = [2]int{r.tail[0] + moveBy[0], r.tail[1] + moveBy[1]}
	r.tailTrace[coordToBinary(r.tail[0], r.tail[1])] = struct{}{}
	r.tailTraceString[coordToString(r.tail[0], r.tail[1])] = struct{}{}

	return nil
}

func (r *rope) placesBeen() int {
	return len(r.tailTrace)
}

func (r *rope) placesBeenString() int {
	return len(r.tailTraceString)
}

func coordToBinary(x, y int) int {
	x += 512
	y += 512
	return x<<10 | y
}

func coordToString(x, y int) string {
	return fmt.Sprintf("%d - %d", x, y)
}

func dirToString(dir int) string {
	switch dir {
	case up:
		return "up"
	case left:
		return "left"
	case down:
		return "down"
	case right:
		return "right"
	default:
		return "what are you thinking"
	}
}
