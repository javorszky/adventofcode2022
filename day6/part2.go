package day6

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

const messageMarkerLength = 14

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 6).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day6/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	stream := gog[0]

	messageMarker, err := extractMarker(stream, messageMarkerLength)
	if err != nil {
		localLogger.Err(err).Msgf("extracting marker failed")
		os.Exit(1)
	}
	// code goes here

	solution := messageMarker
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("-- change this to the part 2 message -- %d", solution)
}
