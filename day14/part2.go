package day14

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 14).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day14/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	r := newRock()
	err = r.addRocksPart2(gog)
	if err != nil {
		localLogger.Err(err).Msgf("adding lines to rock")
		os.Exit(1)
	}

	entryCoord := xyToBinary(0, 500)
	count := 0
	for {
		s := newSand(r.grid)
		restCoord, err := s.findRestingPlace()
		if errors.Is(err, errAbyss) {
			localLogger.Err(err).Msgf("in part 2 sand should not fall to the abyss")
			os.Exit(1)
		}

		r.addSand(restCoord)
		count++

		if entryCoord == restCoord {
			break
		}
	}

	solution := count
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Once all the sand flows to the ground, there have been %d units.", solution)
}
