package day12

import (
	"fmt"
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

const chars = "abcdefghijklmnopqrstuvwxyz SE"

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 12).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day12/input1.txt")
	// gog, err := inputs.ReadIntoLines("day12/input_example.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	em := parseInput(gog, localLogger)

	localLogger.Info().Msgf("The grid\n%s\n\n", em)

	fmt.Printf("start: %v\ngoal: %v\n", em.start, em.goal)

	routes := em.shortestRoute(em.start, func(coord int, _ int32) bool {
		return coord == em.goal
	}, func(currentElevation int32, previousElevation int32) error {
		if currentElevation-previousElevation > 1 {
			return fmt.Errorf("elevation distance too tall: %d -> %d", previousElevation, currentElevation)
		}

		return nil
	})

	for _, r := range routes {
		localLogger.Debug().Msgf("Route length is %d", len(r))

		for _, c := range r {
			row, col := binaryToCoord(c)
			fmt.Printf("%d, %d\n", row, col)
		}
	}

	if len(routes) == 0 {
		localLogger.Error().Msgf("there were no suitable routes found at all!")
	}

	solution := len(routes[0]) - 1
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Shortest route to get to the peak is %d steps", solution)
}
