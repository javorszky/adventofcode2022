package day15

import (
	"os"
	"regexp"
	"strconv"

	"github.com/pkg/errors"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

const lineToCheck = 10

var reInput = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 15).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day15/input_example.txt")
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

		localLogger.Debug().Msgf("adding sensor %v to grid", sensor)
		g.addSensor(sensor)
	}

	sensorsForRow := g.sensorsExcludingRow(lineToCheck)
	localLogger.Debug().Msgf("filtered sensors for row %d: %v", lineToCheck, sensorsForRow)
	lines := make([]line, 0)
	for _, sensorForRow := range sensorsForRow {
		localLogger.Debug().Msgf("\n\nsensor%v", sensorForRow)
		l, err := sensorForRow.lineForRow(lineToCheck)
		if err != nil {
			localLogger.Err(err).Msgf("grabbing bound line for row d", lineToCheck)
			os.Exit(1)
		}

		localLogger.Debug().Msgf("line for sensor for row: %v", l)

		lines = append(lines, l)
	}

	localLogger.Debug().Msgf("before merge: %v", lines)
	merged, err := reduceLines(lines)
	if err != nil {
		localLogger.Err(err).Msgf("reducing lines %v", lines)
		os.Exit(1)
	}
	localLogger.Debug().Msgf("after merge: %v", merged)

	n := 0
	for _, m := range merged {
		n += m.Len()
	}

	solution := n
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("There are %d places where beacon can't be in row %d", solution, lineToCheck)
}

// parseLine takes a line of input, and turns that into a sensor - beacon pair.
// x is row
// y is column
//
// Sensor at x=3056788, y=2626224: closest beacon is at x=3355914, y=2862466
func parseLine(line string) (sensor, error) {
	matches := reInput.FindStringSubmatch(line)
	if len(matches) != 5 {
		return sensor{}, errors.New("match did not happen or happened incorrectly")
	}

	nums := make([]int, 4)
	for i := 1; i < len(matches); i++ {
		n, err := strconv.Atoi(matches[i])
		if err != nil {
			return sensor{}, errors.Wrapf(err, "parsing %s", matches[i])
		}

		nums[i-1] = n
	}

	return newSensor(coordinate{nums[0], nums[1]}, coordinate{nums[2], nums[3]}), nil
}
