package day13

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 13).Int("part", 1).Logger()

	// gog, err := inputs.GroupByBlankLines("day13/input_example.txt")
	gog, err := inputs.GroupByBlankLines("day13/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	pairsInTheCorrectOrder := 0
	for i, pair := range gog {
		left, _, err := parseLine(pair[0], 0)
		if err != nil {
			localLogger.Err(err).Msgf("parsing first line '%s' in group %d", pair[0], i)
			os.Exit(1)
		}

		right, _, err := parseLine(pair[1], 0)
		if err != nil {
			localLogger.Err(err).Msgf("parsing second line '%s' in group %d", pair[1], i)
			os.Exit(1)
		}

		res := smallerList(left, right)
		if res == continueEvaluation {
			localLogger.Error().Msgf("comparing the following two lines resulted in an inconclusive result. This should not have happened!\n%s\n%s", left, right)
		}
		if res == correctOrder {
			pairsInTheCorrectOrder += (i + 1)
		}
	}

	s := localLogger.With().Int("solution", pairsInTheCorrectOrder).Logger()
	s.Info().Msgf("The sum of indices of pairs that are correct is %d", pairsInTheCorrectOrder)
}
