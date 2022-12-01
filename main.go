package main

import (
	"github.com/javorszky/adventofcode2022/day1"

	"github.com/rs/zerolog"
)

func main() {
	l := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Str("module", "adventofcode").Int("year", 2022).Logger()
	l.Info().Msg("Welcome to Gabor Javorszky's Advent of Code 2022 solutions!")

	day1.Task1(l)
	day1.Task2(l)
}
