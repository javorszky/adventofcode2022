package day11

import (
	"github.com/rs/zerolog"
)

// reject modernity, embrace monke

// by.
const (
	// this is 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19, the product of all prime numbers that the monkeys use to test
	// divisibility by.
	primeProduct = 9699690

	// this is 13 * 17 * 19 * 23, the product of all prime numbers that the monkeys in the example use to test
	// divisibility by.
	examplePrimeProduct = 96577
)

type monke struct {
	l        zerolog.Logger
	op       func(int) int
	div      int
	mod      func(int) int
	coolDown func(int) int
	preSend  func(int) int
	success  *monke
	fail     *monke
	items    chan int
	activity int
}

func (m *monke) receive(item int) {
	m.l.Debug().Msgf("received %d, normalizing", item)

	m.items <- item
}

func (m *monke) setSuccessMonke(sm *monke) {
	m.success = sm
}

func (m *monke) setFailMonke(fm *monke) {
	m.fail = fm
}

func (m *monke) run() {
	for {
		if len(m.items) == 0 {
			break
		}
		next := <-m.items
		m.activity++
		inspected := m.op(next)
		cooledDown := m.coolDown(inspected)
		pre := m.preSend(cooledDown)

		if m.mod(pre) == 0 {
			m.l.Debug().Msgf("sent %d to success", pre)

			m.success.receive(pre)
			continue
		}

		m.l.Debug().Msgf("sent %d to fail", pre)
		m.fail.receive(pre)
	}
}

func (m *monke) steps() int {
	return m.activity
}

func newMonke(
	l zerolog.Logger,
	items []int,
	op func(int) int,
	div int,
	coolDown func(int) int,
	preSend func(int) int,
) *monke {
	m := &monke{
		l:        l,
		items:    make(chan int, 40),
		op:       op,
		div:      div,
		mod:      generateModuloFn(div),
		coolDown: coolDown,
		preSend:  preSend,
		activity: 0,
	}

	for _, i := range items {
		m.receive(i)
	}

	return m
}

func getMonkes(l zerolog.Logger, cd func(int) int) []*monke {
	m0 := newMonke(
		l.With().Str("monke", "zero").Logger(),
		[]int{
			96,
			60,
			68,
			91,
			83,
			57,
			85,
		},
		func(i int) int {
			return i * 2
		},
		17,
		cd,
		preSend,
	)

	m1 := newMonke(
		l.With().Str("monke", "one").Logger(),
		[]int{
			75,
			78,
			68,
			81,
			73,
			99,
		},
		func(i int) int {
			return i + 3
		},
		13,
		cd,
		preSend,
	)

	m2 := newMonke(
		l.With().Str("monke", "two").Logger(),
		[]int{
			69,
			86,
			67,
			55,
			96,
			69,
			94,
			85},
		func(i int) int {
			return i + 6
		},
		19,
		cd,
		preSend,
	)

	m3 := newMonke(
		l.With().Str("monke", "three").Logger(),
		[]int{
			88,
			75,
			74,
			98,
			80,
		},
		func(i int) int {
			return i + 5
		},
		7,
		cd,
		preSend,
	)

	m4 := newMonke(
		l.With().Str("monke", "four").Logger(),
		[]int{82},
		func(i int) int {
			return i + 8
		},
		11,
		cd,
		preSend,
	)

	m5 := newMonke(
		l.With().Str("monke", "five").Logger(),
		[]int{
			72,
			92,
			92,
		},
		func(i int) int {
			return i * 5
		},
		3,
		cd,
		preSend,
	)

	m6 := newMonke(
		l.With().Str("monke", "six").Logger(),
		[]int{
			74,
			61,
		},
		func(i int) int {
			return i * i
		},
		2,
		cd,
		preSend,
	)

	m7 := newMonke(
		l.With().Str("monke", "seven").Logger(),
		[]int{
			76,
			86,
			83,
			55,
		},
		func(i int) int {
			return i + 4
		},
		5,
		cd,
		preSend,
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

	return []*monke{m0, m1, m2, m3, m4, m5, m6, m7}
}

func getExampleMonkes(l zerolog.Logger, cd func(int) int) []*monke {
	m0 := newMonke(
		l.With().Str("monke", "zero").Logger(),
		[]int{
			79,
			98,
		},
		func(i int) int {
			return i * 19
		},
		23,
		cd,
		preSendExample,
	)

	m1 := newMonke(
		l.With().Str("monke", "one").Logger(),
		[]int{
			54,
			65,
			75,
			74,
		},
		func(i int) int {
			return i + 6
		},
		19,
		cd,
		preSendExample,
	)

	m2 := newMonke(
		l.With().Str("monke", "two").Logger(),
		[]int{
			79,
			60,
			97,
		},
		func(i int) int {
			return i * i
		},
		13,
		cd,
		preSendExample,
	)

	m3 := newMonke(
		l.With().Str("monke", "three").Logger(),
		[]int{
			74,
		},
		func(i int) int {
			return i + 3
		},
		17,
		cd,
		preSendExample,
	)

	m0.setSuccessMonke(m2)
	m0.setFailMonke(m3)

	m1.setSuccessMonke(m2)
	m1.setFailMonke(m0)

	m2.setSuccessMonke(m1)
	m2.setFailMonke(m3)

	m3.setSuccessMonke(m0)
	m3.setFailMonke(m1)

	return []*monke{m0, m1, m2, m3}
}

func preSend(in int) int {
	return in % primeProduct
}

func preSendExample(in int) int {
	return in & examplePrimeProduct
}
