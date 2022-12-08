package day8

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 8).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day8/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	f, err := newForest(gog, localLogger)
	if err != nil {
		localLogger.Err(err).Msg("trying to create a new forest")
	}

	solution := f.mostScenic()
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("The most scenic score possible for the forest is %d", solution)
}
