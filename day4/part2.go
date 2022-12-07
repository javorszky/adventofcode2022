package day4

import (
	"os"

	"github.com/pkg/errors"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 4).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day4/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	pairs := 0
	// code goes here
	for _, l := range gog {
		cont, err := overlaps(l)
		if err != nil {
			localLogger.Err(err).Msgf("could not calculate fully contains for line %s", l)
			os.Exit(1)
		}

		pairs += cont
	}

	solution := pairs
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("There are %d pairs that have some overlap.", solution)
}

func overlaps(in string) (int, error) {
	mapFirst, mapSecond, err := startEndMaps(in)
	if err != nil {
		return 0, errors.Wrapf(err, "startEndMaps for %s failed", in)
	}

	// find the length of the two maps assuming they do not overlap.
	lenFirst := len(mapFirst)
	lenSecond := len(mapSecond)
	separateLen := lenFirst + lenSecond

	// move elements of the second map into the first map, essentially just merging them together.
	for k := range mapSecond {
		mapFirst[k] = struct{}{}
	}

	// check if the length of the map we merged the other one into is shorter than if the two were separate. In that
	// they overlap somewhat.
	if len(mapFirst) < separateLen {
		return 1, nil
	}

	// if the length did not grow in size, one was fully contained in the other.
	return 0, nil
}
