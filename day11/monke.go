package day11

import "github.com/rs/zerolog"

// reject modernity, embrace monke

type monke struct {
	l        zerolog.Logger
	op       func(int) int
	test     func(int) bool
	success  func(int)
	fail     func(int)
	items    chan int
	activity int
}

func (m *monke) receive(item int) {
	m.l.Debug().Msgf("received %d", item)
	m.items <- item
}

func (m *monke) setSuccessMonke(sm *monke) {
	m.success = sm.receive
}

func (m *monke) setFailMonke(fm *monke) {
	m.fail = fm.receive
}

func (m *monke) coolDown(item int) int {
	return item / 3
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

		if m.test(cooledDown) {
			m.l.Debug().Msgf("sent %d to success", cooledDown)
			m.success(cooledDown)
			continue
		}

		m.l.Debug().Msgf("sent %d to fail", cooledDown)
		m.fail(cooledDown)
	}
}

func (m *monke) holding() []int {
	putBack := make([]int, 0, len(m.items))
	for next := range m.items {
		putBack = append(putBack, next)
	}

	for _, e := range putBack {
		m.receive(e)
	}

	return putBack
}

func (m *monke) steps() int {
	return m.activity
}

func newMonke(l zerolog.Logger, items []int, op func(int) int, test func(int) bool) *monke {
	m := &monke{
		l:        l,
		items:    make(chan int, 40),
		op:       op,
		test:     test,
		activity: 0,
	}

	for _, i := range items {
		m.receive(i)
	}

	return m
}
