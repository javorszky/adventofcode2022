package day15

type grid struct {
	sensors []sensor
}

func (g *grid) addSensor(s sensor) {
	g.sensors = append(g.sensors, s)
}

func (g *grid) sensorsExcludingRow(row int) []sensor {
	listOfSensors := make([]sensor, 0)

	for _, s := range g.sensors {
		if s.rowInExclusion(row) {
			listOfSensors = append(listOfSensors, s)
		}
	}

	return listOfSensors
}

func (g *grid) sensorsBeaconsOnRow(row int) []coordinate {
	out := make([]coordinate, 0)
	for _, sb := range g.sensors {
		if sb.self[1] == row {
			out = append(out, sb.self)
		}

		if sb.closestBeacon[1] == row {
			out = append(out, sb.closestBeacon)
		}
	}

	return uniqueCoordinates(out)
}

func newGrid() grid {
	return grid{sensors: make([]sensor, 0)}
}
