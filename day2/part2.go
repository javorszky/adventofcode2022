package day2

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

// Task2 is functionality to calculate part 2 of the day's task.
//
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won
//
// # A for Rock, B for Paper, and C for Scissors
//
// X means you need to lose
// Y means you need to end the round in a draw
// and Z means you need to win
func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 2).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day2/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	//
	// # A for Rock, B for Paper, and C for Scissors
	//
	// X means you need to lose
	// Y means you need to end the round in a draw
	// and Z means you need to win
	outcomes := map[string]int{
		"A X": 3 + 0, // them rock me lose => scissors
		"A Y": 1 + 3, // them rock me draw => rock
		"A Z": 2 + 6, // them rock me win => paper
		"B X": 1 + 0, // them paper me lose => rock
		"B Y": 2 + 3, // them paper me draw => paper
		"B Z": 3 + 6, // them paper me win => scissors
		"C X": 2 + 0, // them scissors me lose => paper
		"C Y": 3 + 3, // them scissors me draw => scissors
		"C Z": 1 + 6, // them scissors me win => rock
	}

	score := 0
	for i, line := range gog {
		n, ok := outcomes[line]
		if !ok {
			localLogger.Fatal().Msgf("encountered a possibility not in the outcomes on line %d: %s", i, line)
		}
		score += n
	}

	s := localLogger.With().Int("solution", score).Logger()
	s.Info().Msgf("Total score according to strategy guide in part 2 would be %d", score)
}
