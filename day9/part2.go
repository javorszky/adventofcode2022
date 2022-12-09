package day9

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 9).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day9/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	r := newRopeTwo(localLogger)

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

	solution := r.placesBeen()
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("With a 10 link chain, the tail has been to %d places", solution)
}
