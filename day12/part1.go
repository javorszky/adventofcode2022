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
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	em := parseInput(gog)
	fmt.Printf("start: %v\ngoal: %v\n", em.start, em.goal)

	// code goes here

	solution := 2
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("-- change this for the part 1 message -- %d", solution)
}
