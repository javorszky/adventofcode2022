package day8

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 8).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day8/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	f, err := newForest(gog, localLogger)
	if err != nil {
		localLogger.Err(err).Msg("trying to create a new forest")
	}

	solution := f.countVisible()
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("From the edges there are a total of %d trees visible in my forest", solution)
}
