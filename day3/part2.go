package day3

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 3).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day3/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	sum := 0
	aMap, bMap, cMap := make(map[int32]struct{}), make(map[int32]struct{}), make(map[int32]struct{})

	// code goes here
	for i, line := range gog {
		switch i % 3 {
		case 0:
			aMap = make(map[int32]struct{})
			aMap = makeMap(line)
		case 1:
			bMap = make(map[int32]struct{})
			for _, c := range line {
				if _, ok := aMap[c]; ok {
					bMap[c] = struct{}{}
				}
			}
		case 2:
			cMap = make(map[int32]struct{})
			for _, c := range line {
				if _, ok := bMap[c]; ok {
					cMap[c] = struct{}{}
				}
			}

			for k := range cMap {
				prio, err := convertToPriority(k)
				if err != nil {
					localLogger.Err(err).Msgf("could not convert item type %s (%d) to priority: %s", string(k), k, err)
					os.Exit(1)
				}

				sum += prio
			}
		default:
			localLogger.Fatal().Msg("this should never have happened... i %% 3 is not 0, 1, 2")
		}
	}

	solution := sum
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("The group badges sums is %d", solution)
}
