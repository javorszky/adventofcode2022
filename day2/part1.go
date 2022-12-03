package day2

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

// Task1 is the functionality needed to solve part 1 of the day
//
// (1 for Rock, 2 for Paper, and 3 for Scissors)
// (0 if you lost, 3 if the round was a draw, and 6 if you won)
// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors
func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 2).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day2/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	outcomes := map[string]int{
		"A X": 1 + 3, // them rock me rock draw
		"A Y": 2 + 6, // them rock me paper win
		"A Z": 3 + 0, // them rock me scissors lose
		"B X": 1 + 0, // them paper me rock lose
		"B Y": 2 + 3, // them paper me paper draw
		"B Z": 3 + 6, // them paper me scissors win
		"C X": 1 + 6, // them scissors me rock win
		"C Y": 2 + 0, // them scissors me paper lose
		"C Z": 3 + 3, // them scissors me scissors draw
	}

	score := 0
	for i, line := range gog {
		n, ok := outcomes[line]
		if !ok {
			localLogger.Fatal().Msgf("encountered a possibility not in the outcomes on line %d: %s", i, line)
		}
		score += n
	}

	solution := score
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Total score according to strategy guide would be %d", solution)
}
