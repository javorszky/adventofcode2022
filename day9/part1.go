package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 9).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day9/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	r := newRope(localLogger)

	for i, l := range gog {
		dir, dist, err := parseCommand(l)
		if err != nil {
			localLogger.Err(err).Msgf("while parsing string '%s' into command on line %d", l, i)
			os.Exit(1)
		}
		err = r.moveHead(dir, dist)
		if err != nil {
			localLogger.Err(err).Msgf("while moving head to %d %d per command '%s' on line %d", dir, dist, l, i)
			os.Exit(1)
		}
	}

	// code goes here

	solutionStrong := r.placesBeenString()
	solution := r.placesBeen()
	s := localLogger.With().Int("solution", solution).Logger()
	s = s.With().Int("solution string", solutionStrong).Logger()
	s.Info().Msgf("While moving the rope, the tail has been to %d / %d places", solution, solutionStrong)
}

// parseCommand returns the command from the line. Direction, distance, error.
//
// "D 4" would be "down 4", which should be 2,4, nil
func parseCommand(line string) (int, int, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("unexpected command line: %s -- parts: %#v", line, parts)
	}

	dist, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing distance string into int: %s", parts[1])
	}

	switch parts[0] {
	case "U":
		return up, dist, nil
	case "R":
		return right, dist, nil
	case "D":
		return down, dist, nil
	case "L":
		return left, dist, nil
	default:
		return 0, 0, fmt.Errorf("unexpected direction: %s", parts[0])
	}
}
