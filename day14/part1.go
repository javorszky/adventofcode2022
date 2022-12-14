package day14

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 14).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day14/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	r := newRock()
	err = r.addRocks(gog)
	if err != nil {
		localLogger.Err(err).Msgf("adding lines to rock")
		os.Exit(1)
	}

	localLogger.Info().Msgf("The map:\n%s", r.String())

	solution := 2
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("-- change this for the part 1 message -- %d", solution)
}
