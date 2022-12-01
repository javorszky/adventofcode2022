package day1

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/javorszky/adventofcode2022/inputs"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 1).Int("part", 1).Logger()

	gog, err := inputs.GroupByBlankLines("day1/inpdut1_example.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}
	fmt.Printf("%#v", gog)
}
