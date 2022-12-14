package day13

import (
	"fmt"
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 13).Int("part", 1).Logger()

	gog, err := inputs.GroupByBlankLines("day13/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	fmt.Printf("\n\n%v\n\n", gog[0])

	_, _, _ = parseLine("0123456789", 0)
	// code goes here

	solution := 2
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("-- change this for the part 1 message -- %d", solution)
}
