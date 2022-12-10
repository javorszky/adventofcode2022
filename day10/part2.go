package day10

import (
	"os"
	"strings"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

const (
	screenHeight = 6
	screenWidth  = 40
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 10).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day10/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

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

	localLogger.Debug().Msgf("%#v", cycleValues)

	var sb strings.Builder
	for row := 0; row < screenHeight; row++ {
		for col := 0; col < screenWidth; col++ {
			tick := row*screenWidth + col + 1
			regVal := cycleValues[tick]
			localLogger.Debug().Msgf("tick %d, regVal %d, col %d", tick, regVal, col)
			switch {
			case regVal == col || regVal-1 == col || regVal+1 == col:
				localLogger.Debug().Msgf("--- writing #")
				sb.WriteString("#")
			default:
				localLogger.Debug().Msgf("--- writing .")
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}

	solution := sb.String()
	s := localLogger.With().Str("solution", solution).Logger()
	s.Info().Msgf("CRT display ends up being this:\n%s\n", solution)
}
