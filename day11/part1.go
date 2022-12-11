package day11

import (
	"sort"

	"github.com/rs/zerolog"
)

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 11).Int("part", 1).Logger()
	// no file read programmatically

	m0 := newMonke(
		l.With().Str("monke", "zero").Logger(),
		[]int{96, 60, 68, 91, 83, 57, 85},
		func(i int) int {
			return i * 2
		},
		func(i int) bool {
			return i%17 == 0
		},
	)

	m1 := newMonke(
		l.With().Str("monke", "one").Logger(),
		[]int{75, 78, 68, 81, 73, 99},
		func(i int) int {
			return i + 3
		},
		func(i int) bool {
			return i%13 == 0
		},
	)

	m2 := newMonke(
		l.With().Str("monke", "two").Logger(),
		[]int{69, 86, 67, 55, 96, 69, 94, 85},
		func(i int) int {
			return i + 6
		},
		func(i int) bool {
			return i%19 == 0
		},
	)

	m3 := newMonke(
		l.With().Str("monke", "three").Logger(),
		[]int{88, 75, 74, 98, 80},
		func(i int) int {
			return i + 5
		},
		func(i int) bool {
			return i%7 == 0
		},
	)

	m4 := newMonke(
		l.With().Str("monke", "four").Logger(),
		[]int{82},
		func(i int) int {
			return i + 8
		},
		func(i int) bool {
			return i%11 == 0
		},
	)

	m5 := newMonke(
		l.With().Str("monke", "five").Logger(),
		[]int{72, 92, 92},
		func(i int) int {
			return i * 5
		},
		func(i int) bool {
			return i%3 == 0
		},
	)

	m6 := newMonke(
		l.With().Str("monke", "six").Logger(),
		[]int{74, 61},
		func(i int) int {
			return i * i
		},
		func(i int) bool {
			return i%2 == 0
		},
	)

	m7 := newMonke(
		l.With().Str("monke", "seven").Logger(),
		[]int{76, 86, 83, 55},
		func(i int) int {
			return i + 4
		},
		func(i int) bool {
			return i%5 == 0
		},
	)

	m0.setSuccessMonke(m2)
	m0.setFailMonke(m5)

	m1.setSuccessMonke(m7)
	m1.setFailMonke(m4)

	m2.setSuccessMonke(m6)
	m2.setFailMonke(m5)

	m3.setSuccessMonke(m7)
	m3.setFailMonke(m1)

	m4.setSuccessMonke(m0)
	m4.setFailMonke(m2)

	m5.setSuccessMonke(m6)
	m5.setFailMonke(m3)

	m6.setSuccessMonke(m3)
	m6.setFailMonke(m1)

	m7.setSuccessMonke(m4)
	m7.setFailMonke(m0)

	round := []*monke{m0, m1, m2, m3, m4, m5, m6, m7}

	localLogger.Debug().Msg("all right, let's kick it off!")

	for i := 0; i < 20; i++ {
		for _, currentMonke := range round {
			currentMonke.run()
		}
	}

	localLogger.Debug().Msg("time to count stuff")

	sort.Slice(round, func(i, j int) bool {
		return round[i].steps() > round[j].steps()
	})

	localLogger.Debug().Msgf("and now for the final")

	solution := round[0].steps() * round[1].steps()
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("Monke business for part 1 after 20 rounds is %d", solution)
}
