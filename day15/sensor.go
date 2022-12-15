package day15

import (
	"fmt"
)

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

func (s sensor) rowBoundCoordinates(row int) (coordinate, coordinate, error) {
	if !s.rowInExclusion(row) {
		return coordinate{}, coordinate{}, fmt.Errorf("row %d is not in exclusion zone for this sensor", row)
	}

	d := s.manhattanDistance - absDiff(row, s.self[1])

	return coordinate{s.self[0] - d, row}, coordinate{s.self[0] + d, row}, nil
}

type grid struct {
	sensors []sensor
}

func (g grid) addSensor(s sensor) {
	g.sensors = append(g.sensors, s)
}

func (g grid) sensorsExcludingRow(row int) []sensor {
	listOfSensors := make([]sensor, 0)

	for _, s := range g.sensors {
		if s.rowInExclusion(row) {
			listOfSensors = append(listOfSensors, s)
		}
	}

	return listOfSensors
}

func newSensor(own, closest coordinate) sensor {
	return sensor{
		self:              own,
		closestBeacon:     closest,
		manhattanDistance: manhattanDistance(own, closest),
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

func manhattanDistance(a, b coordinate) int {
	return absDiff(a[0], b[0]) + absDiff(a[1], b[1])
}
