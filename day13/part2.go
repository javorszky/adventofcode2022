package day13

import (
	"os"
	"sort"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

const (
	divider2 = "[[2]]"
	divider6 = "[[6]]"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 13).Int("part", 2).Logger()

	gog, err := inputs.GroupByBlankLines("day13/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	pairs := make([]list, len(gog)*2)
	for i, group := range gog {
		left, _, err := parseLine(group[0], 0)
		if err != nil {
			localLogger.Err(err).Msgf("parsing first line line '%s' from group %d", group[0], i)
			os.Exit(1)
		}
		right, _, err := parseLine(group[1], 0)
		if err != nil {
			localLogger.Err(err).Msgf("parsing second line line '%s' from group %d", group[1], i)
			os.Exit(1)
		}
		pairs[i*2] = left
		pairs[i*2+1] = right
	}

	d2, _, err := parseLine(divider2, 0)
	if err != nil {
		localLogger.Err(err).Msgf("parsing divider 2 line '%s'", divider2)
		os.Exit(1)
	}

	d6, _, err := parseLine(divider6, 0)
	if err != nil {
		localLogger.Err(err).Msgf("parsing divider 6 line '%s'", divider6)
		os.Exit(1)
	}

	pairs = append(pairs, d2, d6)

	sort.Slice(pairs, func(i, j int) bool {
		res := smallerList(pairs[i], pairs[j])
		return res == correctOrder
	})

	product := 1
	for i, l := range pairs {
		if l.String() == divider2 || l.String() == divider6 {
			product = product * (i + 1)
		}
	}

	solution := product
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("After dividers are put in and lines are sorted, the product of the dividers' indices comes to %d", solution)
}
