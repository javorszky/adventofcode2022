package day5

import (
	"os"
	"strings"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 5).Int("part", 1).Logger()

	gog, err := inputs.GroupByBlankLines("day5/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	st := makeStacks(gog[0])

	insts, err := makeInstructions(gog[1])
	if err != nil {
		localLogger.Err(err).Msgf("parsing instructions")
		os.Exit(1)
	}

	for _, inst := range insts {
		for i := 0; i < inst[0]; i++ {
			st[inst[2]].Push(st[inst[1]].Pop())
		}
	}

	var sb strings.Builder
	for i := 1; i < 10; i++ {
		sb.WriteString(st[i].Pop())
	}
	// code goes here

	solution := sb.String()
	s := localLogger.With().Str("solution", solution).Logger()
	s.Info().Msgf("After shuffling the boxes, the top ones are '%s'", solution)
}
