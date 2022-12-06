package day6

import (
	"os"
	"sort"
	"strings"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 6).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day6/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	stream := gog[0]

	bla, err := extractMarker(stream)
	if err != nil {
		localLogger.Err(err).Msgf("extracting marker failed")
		os.Exit(1)
	}
	// code goes here

	solution := bla
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Offset for the first unique 4 char sequence is %d", solution)
}

func extractMarker(in string) (int, error) {
	sbuf := strings.NewReader(in)

	marker := make([]byte, 4)
	var off int64 = 0
	for {
		at, err := sbuf.ReadAt(marker, off)
		if err != nil || at < 4 {
			return 0, err
		}

		sort.Slice(marker, func(i, j int) bool {
			return marker[i] > marker[j]
		})
		compactedMarker := slices.Compact(marker)
		if len(marker) == len(compactedMarker) {
			return int(off) + 4, nil
		}

		off++
	}
}
