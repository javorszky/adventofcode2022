package day12

import (
	"fmt"
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 12).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day12/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	em := parseInput(gog, localLogger)

	localLogger.Info().Msgf("The grid\n%s\n\n", em)

	routes := em.shortestRoute(
		em.goal,
		func(coord int, el int32) bool {
			v, ok := em.grid[coord]
			if !ok {
				return false
			}
			return v == 97
		},
		func(currentElevation int32, previousElevation int32) error {
			// previous is big
			// previous can be at most 1 bigger
			// previous can be 20 smaller

			// previous 100
			// current 80  that's not
			//
			// previous 100
			// current 102 good

			if previousElevation-currentElevation > 1 {
				return fmt.Errorf("can't ascent: %d -> %d", previousElevation, currentElevation)
			}

			return nil
		},
	)

	if len(routes) == 0 {
		localLogger.Error().Msgf("there were no suitable routes found at all!")
	}

	solution := len(routes[0]) - 1
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Shortest hike from an 'a' elevation is %d steps", solution)
}
