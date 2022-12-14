package day14

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/pkg/errors"
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

	count := 0
	for {
		s := newSand(r.grid)
		restCoord, err := s.findRestingPlace()
		if errors.Is(err, errAbyss) {
			break
		}
		r.addSand(restCoord)
		count++
	}

	solution := count
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("At the end %d units of sand came to rest before abyss", solution)
}
