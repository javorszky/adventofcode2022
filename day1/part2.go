package day1

import (
	"os"
	"sort"

	"github.com/rs/zerolog"

	"github.com/javorszky/adventofcode2022/inputs"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 1).Int("part", 2).Logger()

	gog, err := inputs.GroupByBlankLines("day1/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	gogi, err := convertToInts(gog)
	if err != nil {
		localLogger.Err(err).Msg("converting group of lines to group of ints")
		os.Exit(1)
	}

	rolledUp := make([]int, len(gogi))
	for i, g := range gogi {
		rolledUp[i] = reduceInts(g)
	}

	sort.Ints(rolledUp)

	topThree := []int{
		rolledUp[len(rolledUp)-1],
		rolledUp[len(rolledUp)-2],
		rolledUp[len(rolledUp)-3],
	}

	solution := reduceInts(topThree)

	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("The three elves carrying the most calories are carrying a total of %d", solution)
}
