package day11

import (
	"sort"

	"github.com/rs/zerolog"
)

func Task2(l zerolog.Logger) {
	localLogger := l.With().Int("day", 11).Int("part", 2).Logger()

	cd := func(item int) int {
		return item
	}

	//mks := getMonkes(localLogger, cd)
	mks := getExampleMonkes(localLogger, cd)

	localLogger.Debug().Msg("all right, let's kick it off!")

	for i := 0; i < 1000; i++ {
		localLogger.Info().Msgf("iter %d", i)
		for _, currentMonke := range mks {
			currentMonke.run()
		}
	}

	localLogger.Debug().Msg("time to count stuff")

	sort.Slice(mks, func(i, j int) bool {
		return mks[i].steps() > mks[j].steps()
	})

	localLogger.Info().Msgf("%#v", mks)

	localLogger.Debug().Msgf("and now for the final")

	solution := mks[0].steps() * mks[1].steps()
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("With no worry cooldown and 10,000 rounds, the product of the top two activity monkey is %d", solution)
}
