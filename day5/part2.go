package day5

import (
	"os"
	"strings"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 5).Int("part", 2).Logger()

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
		els, err := st[inst[1]].PopMany(inst[0])
		if err != nil {
			localLogger.Err(err).Msgf("move %d from %d to %d", inst[0], inst[1], inst[2])
		}
		st[inst[2]].PushMany(els)
	}

	var sb strings.Builder
	for i := 1; i < 10; i++ {
		sb.WriteString(st[i].Pop())
	}
	// code goes here

	solution := sb.String()

	s := localLogger.With().Str("solution", solution).Logger()
	s.Info().Msgf("When using the CrateMover 9001 with its bulk functionality, the top crates are '%s'", solution)
}
