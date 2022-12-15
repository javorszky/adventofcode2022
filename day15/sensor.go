package day15

// coordinate holds a pair of ints, in col / row format.
// x, y =>
//
//	x is column, horizontal, right is plus
//	y is row, vertical, down is plus
type coordinate [2]int

type sensor struct {
	self              coordinate
	closestBeacon     coordinate
	manhattanDistance int
}

func (s sensor) rowInExclusion(row int) bool {
	return s.self[1]-s.manhattanDistance <= row && row <= s.self[1]+s.manhattanDistance
}

type grid struct {
	sensors []sensor
}

func (g grid) addSensor(s sensor) {
	g.sensors = append(g.sensors, s)
}

func newSensor(own, closest coordinate) sensor {
	return sensor{
		self:              own,
		closestBeacon:     closest,
		manhattanDistance: absDiff(own[0], closest[0]) + absDiff(own[1], closest[1]),
	}
}

func newGrid() grid {
	return grid{sensors: make([]sensor, 0)}
}

func absDiff(a, b int) int {
	c := a - b
	if c < 0 {
		return b - a
	}
	return c
}
