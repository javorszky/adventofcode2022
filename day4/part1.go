package day4

import (
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 4).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day4/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	pairs := 0
	// code goes here
	for _, l := range gog {
		cont, err := fullyContains(l)
		if err != nil {
			localLogger.Err(err).Msgf("could not calculate fully contains for line %s", l)
			os.Exit(1)
		}

		pairs += cont
	}

	solution := pairs
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("There are %d pairs where one fully contains the other's assignment", solution)
}

func startEndMaps(in string) (map[int]struct{}, map[int]struct{}, error) {
	// split the line into the two groups
	//
	// 1-2,3-4 => 1,2 and 3,4
	pairs := strings.Split(in, ",")
	if len(pairs) != 2 {
		return nil, nil, errors.New("splitting line by comma yielded unexpected number of parts")
	}

	// split the first group into beginning and end numbers
	//
	// 1-2 => 1 and 2
	first := strings.Split(pairs[0], "-")
	if len(first) != 2 {
		return nil, nil, errors.New("splitting first elf assignment by dash yielded unexpected number of parts")
	}

	f1, err := strconv.Atoi(first[0])
	if err != nil {
		return nil, nil, errors.Wrap(err, "converting starting schedule for first elf")
	}
	f2, err := strconv.Atoi(first[1])
	if err != nil {
		return nil, nil, errors.Wrap(err, "converting ending schedule for first elf")
	}

	// split the second group into beginning and end numbers
	//
	// 3-4 => 3 and 4
	second := strings.Split(pairs[1], "-")
	if len(second) != 2 {
		return nil, nil, errors.New("splitting second elf assignment by dash yielded unexpected number of parts")
	}

	s1, err := strconv.Atoi(second[0])
	if err != nil {
		return nil, nil, errors.Wrap(err, "converting starting schedule for second elf")
	}
	s2, err := strconv.Atoi(second[1])
	if err != nil {
		return nil, nil, errors.Wrap(err, "converting ending schedule for second elf")
	}

	// create maps to keep track of overlaps
	mapFirst := make(map[int]struct{})
	mapSecond := make(map[int]struct{})

	// assign first and second beginning numbers to map
	//
	// 1 -> 2 =>
	// map{
	//   1: struct{},
	//   2: struct{},
	// }
	for i := f1; i <= f2; i++ {
		mapFirst[i] = struct{}{}
	}

	for i := s1; i <= s2; i++ {
		mapSecond[i] = struct{}{}
	}

	return mapFirst, mapSecond, nil
}

func fullyContains(one string) (int, error) {
	mapFirst, mapSecond, err := startEndMaps(one)
	if err != nil {
		return 0, errors.Wrapf(err, "startEndMaps for %s failed", one)
	}

	// find the map that's longer. If we're fully enclosing then the longer one needs to fully contain the shorter one.
	// save the longer length in a variable maxLen.
	lenFirst := len(mapFirst)
	lenSecond := len(mapSecond)
	maxLen := lenFirst
	if lenSecond > maxLen {
		maxLen = lenSecond
	}

	// move elements of the second map into the first map, essentially just merging them together.
	for k := range mapSecond {
		mapFirst[k] = struct{}{}
	}

	// check if the length of the map we merged the other one into grew in size. If it did grow in size, one of them
	// wasn't contained in the other one.
	if len(mapFirst) > maxLen {
		return 0, nil
	}

	// if the length did not grow in size, one was fully contained in the other.
	return 1, nil
}
