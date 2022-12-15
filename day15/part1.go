package day15

import (
	"os"
	"regexp"
	"strconv"

	"github.com/pkg/errors"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

var reInput = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 15).Int("part", 1).Logger()

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

		g.addSensor(sensor)
	}

	solution := 2
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("-- change this for the part 1 message -- %d", solution)
}

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
