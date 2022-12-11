package day11

import (
	"sort"

	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 11).Int("part", 1).Logger()
	// no file read programmatically

	round := getMonkes(localLogger, coolDownP1)

	localLogger.Debug().Msg("all right, let's kick it off!")

	for i := 0; i < 20; i++ {
		localLogger.Debug().Msgf("\n\nRound %d\n\n", i)
		for _, currentMonke := range round {
			currentMonke.run()
		}
	}

	localLogger.Debug().Msg("time to count stuff")

	sort.Slice(round, func(i, j int) bool {
		return round[i].steps() > round[j].steps()
	})

	localLogger.Debug().Msgf("and now for the final")

	for _, m := range round {
		localLogger.Debug().Msgf("activity: %d", m.steps())
	}

	solution := round[0].steps() * round[1].steps()
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Monke business for part 1 after 20 rounds is %d", solution)
}

func coolDownP1(in int) int {
	return in / 3
}

func generateModuloFn(in int) func(int) int {
	return func(d int) int {
		return d % in
	}
}
