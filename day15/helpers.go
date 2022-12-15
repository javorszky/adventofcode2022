package day15

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

func pluck[T any](sl []T, idx int) (T, []T, error) {
	var thing T
	if idx >= len(sl) {
		return thing, nil, errors.New("index out of bounds")
	}

	thing = sl[idx]
	ls := append(sl[:idx], sl[idx+1:]...)
	return thing, ls, nil
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

func uniqueCoordinates(in []coordinate) []coordinate {
	m := make(map[string]coordinate)

	for _, c := range in {
		m[fmt.Sprintf("%d,%d", c[0], c[1])] = c
	}

	out := make([]coordinate, len(m))
	i := 0
	for _, v := range m {
		out[i] = v
		i++
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i][0] != out[j][0] {
			return out[i][0] < out[j][0]
		}

		return out[i][1] < out[j][1]
	})

	return out
}
