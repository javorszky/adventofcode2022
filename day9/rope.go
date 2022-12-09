package day9

import (
	"fmt"
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
	head      [2]int
	tail      [2]int
	tailTrace map[uint16]struct{}
	bounds    [4]int
}

func newRope() *rope {
	return &rope{
		head: [2]int{0, 0},
		tail: [2]int{0, 0},
		tailTrace: map[uint16]struct{}{
			0: {},
		},
		bounds: [4]int{0, 0, 0, 0},
	}
}

func (r *rope) moveHead(dir int, dist int) error {
	switch dir {
	case up:
		r.head = [2]int{r.head[0] + dist, r.head[1]}
	case right:
		r.head = [2]int{r.head[0], r.head[1] + dist}
	case down:
		r.head = [2]int{r.head[0] - dist, r.head[1]}
	case left:
		r.head = [2]int{r.head[0], r.head[1] - dist}
	default:
		return fmt.Errorf("this should not have happened, got direction %#v", dir)
	}

	// check if we expanded the up bound
	if r.head[0] > r.bounds[0] {
		r.bounds[0] = r.head[0]
	}

	// check if we expanded the down bound
	if r.head[0] < r.bounds[2] {
		r.bounds[2] = r.head[0]
	}

	// check if we expanded the right bound
	if r.head[1] > r.bounds[1] {
		r.bounds[1] = r.head[1]
	}

	// check if we expanded the left bound
	if r.head[1] < r.bounds[3] {
		r.bounds[3] = r.head[1]
	}

	return nil
}
