package day7

import (
	"os"
	"sort"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	availableDiskSpace = 70000000
	neededDiskSpace    = 30000000
	maxUsedSpace       = availableDiskSpace - neededDiskSpace
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 7).Int("part", 2).Logger()

	gog, err := inputs.ReadIntoLines("day7/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	root := buildDirectory(gog, localLogger)

	candidates := filterDirsAtLeast(root, root.size()-maxUsedSpace)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].size() < candidates[j].size()
	})

	solution := candidates[0].size()
	s := localLogger.With().Int("solution", solution).Logger()
	p := message.NewPrinter(language.English)
	s.Info().Msgf("The size of the smallest directory I can delete to get back the space is %s / %d", p.Sprintf("%d", solution), solution)
}
