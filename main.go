package main

import (
	"os"

	"github.com/javorszky/adventofcode2022/day1"

	"github.com/rs/zerolog"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	l := zerolog.New(os.Stderr).With().Str("module", "adventofcode").Int("year", 2022).Logger()
	l.Info().Msg("Welcome to Gabor Javorszky's Advent of Code 2022 solutions!")

	day1.Task1(l)
}
