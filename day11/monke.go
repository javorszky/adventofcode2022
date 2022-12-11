package day11

import (
	"math/big"

	"github.com/rs/zerolog"
)

// reject modernity, embrace monke

type monke struct {
	l        zerolog.Logger
	op       func(*big.Int) *big.Int
	div      int
	mod      func(*big.Int) *big.Int
	coolDown func(*big.Int) *big.Int
	success  func(*big.Int)
	fail     func(*big.Int)
	items    chan *big.Int
	activity int
}

func (m *monke) receive(item *big.Int) {
	m.l.Debug().Msgf("received %d, normalizing", item)
	bla := m.coolDown(m.op(item))
	mod := m.mod(bla)

	m.items <- item
}

func (m *monke) setSuccessMonke(sm *monke) {
	m.success = sm.receive
}

func (m *monke) setFailMonke(fm *monke) {
	m.fail = fm.receive
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

		if m.mod(cooledDown).Int64() == 0 {
			m.l.Debug().Msgf("sent %d to success", cooledDown)
			m.success(cooledDown)
			continue
		}

		m.l.Debug().Msgf("sent %d to fail", cooledDown)
		m.fail(cooledDown)
	}
}

func (m *monke) steps() int {
	return m.activity
}

func newMonke(l zerolog.Logger, items []*big.Int, op func(*big.Int) *big.Int, div int, coolDown func(*big.Int) *big.Int) *monke {
	m := &monke{
		l:        l,
		items:    make(chan *big.Int, 40),
		op:       op,
		div:      div,
		mod:      generateModuloFn(div),
		coolDown: coolDown,
		activity: 0,
	}

	for _, i := range items {
		m.receive(i)
	}

	return m
}

func getMonkes(l zerolog.Logger, cd func(*big.Int) *big.Int) []*monke {
	m0 := newMonke(
		l.With().Str("monke", "zero").Logger(),
		[]*big.Int{
			big.NewInt(96),
			big.NewInt(60),
			big.NewInt(68),
			big.NewInt(91),
			big.NewInt(83),
			big.NewInt(57),
			big.NewInt(85),
		},
		func(i *big.Int) *big.Int {
			return i.Mul(i, big.NewInt(2))
		},
		17,
		cd,
	)

	m1 := newMonke(
		l.With().Str("monke", "one").Logger(),
		[]*big.Int{
			big.NewInt(75),
			big.NewInt(78),
			big.NewInt(68),
			big.NewInt(81),
			big.NewInt(73),
			big.NewInt(99),
		},
		func(i *big.Int) *big.Int {
			return i.Add(i, big.NewInt(3))
		},
		13,
		cd,
	)

	m2 := newMonke(
		l.With().Str("monke", "two").Logger(),
		[]*big.Int{
			big.NewInt(69),
			big.NewInt(86),
			big.NewInt(67),
			big.NewInt(55),
			big.NewInt(96),
			big.NewInt(69),
			big.NewInt(94),
			big.NewInt(85)},
		func(i *big.Int) *big.Int {
			return i.Add(i, big.NewInt(6))
		},
		19,
		cd,
	)

	m3 := newMonke(
		l.With().Str("monke", "three").Logger(),
		[]*big.Int{
			big.NewInt(88),
			big.NewInt(75),
			big.NewInt(74),
			big.NewInt(98),
			big.NewInt(80),
		},
		func(i *big.Int) *big.Int {
			return i.Add(i, big.NewInt(5))
		},
		7,
		cd,
	)

	m4 := newMonke(
		l.With().Str("monke", "four").Logger(),
		[]*big.Int{big.NewInt(82)},
		func(i *big.Int) *big.Int {
			return i.Add(i, big.NewInt(8))
		},
		11,
		cd,
	)

	m5 := newMonke(
		l.With().Str("monke", "five").Logger(),
		[]*big.Int{
			big.NewInt(72),
			big.NewInt(92),
			big.NewInt(92),
		},
		func(i *big.Int) *big.Int {
			return i.Mul(i, big.NewInt(5))
		},
		3,
		cd,
	)

	m6 := newMonke(
		l.With().Str("monke", "six").Logger(),
		[]*big.Int{
			big.NewInt(74),
			big.NewInt(61),
		},
		func(i *big.Int) *big.Int {
			return i.Mul(i, i)
		},
		2,
		cd,
	)

	m7 := newMonke(
		l.With().Str("monke", "seven").Logger(),
		[]*big.Int{
			big.NewInt(76),
			big.NewInt(86),
			big.NewInt(83),
			big.NewInt(55),
		},
		func(i *big.Int) *big.Int {
			return i.Add(i, big.NewInt(4))
		},
		5,
		cd,
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

func getExampleMonkes(l zerolog.Logger, cd func(*big.Int) *big.Int) []*monke {
	m0 := newMonke(
		l.With().Str("monke", "zero").Logger(),
		[]*big.Int{
			big.NewInt(79),
			big.NewInt(98),
		},
		func(i *big.Int) *big.Int {
			return i.Mul(i, big.NewInt(19))
		},
		23,
		cd,
	)

	m1 := newMonke(
		l.With().Str("monke", "one").Logger(),
		[]*big.Int{
			big.NewInt(54),
			big.NewInt(65),
			big.NewInt(75),
			big.NewInt(74),
		},
		func(i *big.Int) *big.Int {
			return i.Add(i, big.NewInt(6))
		},
		19,
		cd,
	)

	m2 := newMonke(
		l.With().Str("monke", "two").Logger(),
		[]*big.Int{
			big.NewInt(79),
			big.NewInt(60),
			big.NewInt(97),
		},
		func(i *big.Int) *big.Int {
			return i.Mul(i, i)
		},
		13,
		cd,
	)

	m3 := newMonke(
		l.With().Str("monke", "three").Logger(),
		[]*big.Int{
			big.NewInt(74),
		},
		func(i *big.Int) *big.Int {
			return i.Add(i, big.NewInt(3))
		},
		17,
		cd,
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
