package day15

import (
	"os"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

const (
	minV = 0
	maxV = 4000000
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 15).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day15/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	g := newGrid()

	for _, l := range gog {
		sensor, err := parseLine(l)
		if err != nil {
			localLogger.Err(err).Msgf("parsing this line did not match results:\n%s", l)
		}

		// localLogger.Debug().Msgf("adding sensor %v to grid", sensor)
		g.addSensor(sensor)
	}
	part2LineToCheck := 0

	for i := 0; i <= maxV; i++ {
		sensorsForRow := g.sensorsExcludingRow(part2LineToCheck)
		lines := make([]line, 0)
		for _, sensorForRow := range sensorsForRow {
			l, err := sensorForRow.clippedLineForRow(part2LineToCheck)
			if err != nil {
				localLogger.Err(err).Msgf("grabbing bound line for row %d", part2LineToCheck)
				os.Exit(1)
			}

			lines = append(lines, l)
		}

		merged, err := reduceLines(lines)
		if err != nil {
			localLogger.Err(err).Msgf("reducing lines %v", lines)
			os.Exit(1)
		}

		n := 0
		for _, m := range merged {
			n += m.Len()
		}

		if maxV+1-n > 0 {
			localLogger.Debug().Msgf("ROW %02d: merged lines: %v", part2LineToCheck, merged)
		}

		part2LineToCheck++
	}

	// this was manual, see note below
	solution := 11374534948438
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Above there should be a single line with a row number. Find the first line's end, grab the first coordinate, add 1 to it, multiply by 4,000,000, add the other coordinate. In this case the lines are\n\n%s\n\nFrom there the end of the first one is %s, so the (x+1)*4000000+y ends up being %d", "[{[0 2948438] [2843632 2948438] 1 2948438} {[2843634 2948438] [4000000 2948438] 1 2948438}]", "[2843632 2948438]", solution)
}
