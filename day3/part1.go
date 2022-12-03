package day3

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 3).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day3/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	sum := 0

	// code goes here
	for _, line := range gog {
		a, b, err := split(line)
		if err != nil {
			log.Fatalf("could not split line '%s' into two: %s", line, err)
		}

		aMap := makeMap(a)
		common := make([]int32, 0)
		for _, c := range b {
			if _, ok := aMap[c]; ok {
				common = append(common, c)
			}
		}

		common = slices.Compact(common)

		for _, t := range common {
			prio, err := convertToPriority(t)

			if err != nil {
				log.Fatalf("could not convert item type %s (%d) to priority: %s", string(t), t, err)
			}

			sum += prio
		}
	}

	solution := sum
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("The sum of all misplaced items is %d", solution)
}

// split breaks a string into two halves.
func split(in string) (string, string, error) {
	if in == "" || len(in)%2 != 0 {
		return "", "", errors.New("input string is odd length")
	}

	return in[:len(in)/2], in[len(in)/2:], nil
}

// makeMap creates a map out of a string, so we can look up individual characters easily.
func makeMap(in string) map[int32]struct{} {
	m := make(map[int32]struct{})

	for _, c := range in {
		m[c] = struct{}{}
	}

	return m
}

// convertToPriority will take the code point of a character, and return the priority assigned to it by the task.
//
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
//
// Code point ranges:
// a: 97
// z: 122
// A: 65
// Z: 90
func convertToPriority(in int32) (int, error) {
	switch {
	case in >= 96 && in <= 122:
		return int(in - 96), nil
	// it's a lowercase a-z
	case in >= 65 && in <= 90:
		return int(in - 38), nil
	// it's an uppercase A-Z
	default:
		return 0, fmt.Errorf("received item type out of bounds: %s (%d)", string(in), in)
	}
}
