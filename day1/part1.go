package day1

import (
	"os"
	"sort"
	"strconv"

	"github.com/pkg/errors"

	"github.com/rs/zerolog"

	"github.com/javorszky/adventofcode2022/inputs"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 1).Int("part", 1).Logger()

	gog, err := inputs.GroupByBlankLines("day1/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	gogi, err := convertToInts(gog)
	if err != nil {
		localLogger.Err(err).Msg("converting group of lines to group of ints")
		os.Exit(1)
	}

	rolledUp := make([]int, len(gogi))
	for i, g := range gogi {
		rolledUp[i] = reduceInts(g)
	}

	sort.Ints(rolledUp)
	solution := rolledUp[len(rolledUp)-1]

	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("The elf carrying the most calories is carrying %d", solution)
}

func convertToInts(in [][]string) ([][]int, error) {
	out := make([][]int, len(in))
	for i, g := range in {
		temp := make([]int, len(g))
		for j, s := range g {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, errors.Wrapf(err, "converting %s to int", s)
			}
			temp[j] = n
		}
		out[i] = temp
	}

	return out, nil
}

func reduceInts(in []int) int {
	start := 0
	for _, v := range in {
		start += v
	}
	return start
}
