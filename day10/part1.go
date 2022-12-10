package day10

import (
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 10).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day10/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	// 20th, 60th, 100th, 140th, 180th, and 220th
	cycleValues := make([]int, 241)
	cycleValues[0] = 1

	currentCycle := 0
	currentRegister := 1
	for _, line := range gog {
		cycs, rise, err := parseCommand(line)
		if err != nil {
			localLogger.Err(err).Msgf("something went wrong with turning input to command")
			os.Exit(1)
		}
		currentCycle++
		cycleValues[currentCycle] = currentRegister
		if cycs == 2 {
			currentCycle++
			cycleValues[currentCycle] = currentRegister
		}
		currentRegister += rise
	}

	cycles := [6]int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, n := range cycles {
		sum += cycleValues[n] * n
	}

	// code goes here

	solution := sum
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("The sum of the signal strengths is %d", solution)
}

// parseCommand takes a line of the input, and spits out two details, as well as an error in case something went wrong.
//
// The return values are how many cycles the command takes up, and how much the register needs to increase.
func parseCommand(line string) (int, int, error) {
	if line == "noop" {
		return 1, 0, nil
	}

	num := strings.TrimPrefix(line, "addx ")
	nint, err := strconv.Atoi(num)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "line '%s', converting '%s' to int", line, num)
	}

	return 2, nint, nil
}
